package app

import (
	"fmt"
	"net/http"
	"transaction-service/internal/common"
	"transaction-service/internal/infrastructure"

	bun "github.com/uptrace/bunrouter"
)

const (
	LiveZEndpoint                         = "/livez"
	ReadyZEndpoint                        = "/readyz"
	ContentTypeHeaderKey                  = "Content-Type"
	ContentTypeApplicationJsonHeaderValue = "application/json"
	V1BaseEndpoint                        = "/v1"
	UserPathGroup                         = "/users"
)

var (
	transactionEndpoint = fmt.Sprintf("/:%s/transactions", common.UserIDPathParam)
)

func (a *app) setHealthRoute() {
	//* Both of these endpoints make a ping to all dbs configured in the 'repository' layer.
	//* Modify that according to the number of dbs configured in the service, created from this template
	a.router.GET(LiveZEndpoint, func(w http.ResponseWriter, req bun.Request) error {
		// use request context inside all controllers to trace per URL and method
		sLivez := infrastructure.CreateHTTPSpan(req.Context(), "livez")
		// always end spans
		defer sLivez.End()

		// use tracer context to execute queries (this allows the tracer to calculate query time from the request)
		err := a.controllers.HealthCheckController().PingDB(sLivez.Context())
		if err != nil {
			a.log.Errorf("can't ping database(s) %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return bun.JSON(w, bun.H{"message": "error pinging database(s)"})
		}
		return bun.JSON(w, bun.H{"status": "live"})
	})

	a.router.GET(ReadyZEndpoint, func(w http.ResponseWriter, req bun.Request) error {
		sReadyz := infrastructure.CreateHTTPSpan(req.Context(), "livez")
		defer sReadyz.End()

		err := a.controllers.HealthCheckController().PingDB(sReadyz.Context())
		if err != nil {
			a.log.Errorf("can't ping database(s) %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return bun.JSON(w, bun.H{"message": "error pinging database(s)"})
		}
		return bun.JSON(w, bun.H{"status": "ready"})
	})
}

func (a *app) setAPIRoutes() {
	// Transaction group middleware
	v1Group := a.router.NewGroup(V1BaseEndpoint).Use(func(next bun.HandlerFunc) bun.HandlerFunc {
		return func(w http.ResponseWriter, req bun.Request) error {
			w.Header().Add(ContentTypeHeaderKey, ContentTypeApplicationJsonHeaderValue)
			return next(w, req)
		}
	})

	// group for 'users' routes
	v1Group.WithGroup(UserPathGroup, func(g *bun.Group) {
		// mutes/bans user from chat
		g.POST(transactionEndpoint, a.controllers.UserTransactionController().HandleUserTransactionRequest)
	})

}
