CREATE TABLE public.car (
                            id uuid NOT NULL,
                            make character varying(100) NOT NULL,
                            model character varying(100) NOT NULL,
                            year integer NOT NULL,
                            color character varying(50),
                            vin character varying(17) NOT NULL,
                            mileage integer DEFAULT 0 NOT NULL,
                            price numeric(12,2) NOT NULL,
                            disabled boolean DEFAULT false NOT NULL,
                            created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
                            updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.car OWNER TO postgres;



INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440000', 'Toyota', 'Corolla', 2020, 'Red', 'JTDBL40E799123456', 45000, 18999.99, false, '2026-04-15 17:56:31.26811-06', '2026-04-15 17:56:31.26811-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('f6d4daa6-d68f-487a-bc7f-ef31a8dd7747', 'Toyota', 'Rav4', 2022, 'White', 'JTDBL40E799123458', 15000, 18999.99, false, '0001-12-31 17:23:24-06:36:36 BC', '0001-12-31 17:23:24-06:36:36 BC');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('04c509aa-7038-4b6b-bf51-c6167118188a', 'Honda', 'Civic', 2023, 'Blue and Red', '2HGFE2F59PH123457', 8000, 21950.00, false, '0001-12-31 17:23:24-06:36:36 BC', '2026-04-15 19:01:50.566515-06');