--
-- PostgreSQL database dump
--

\restrict xwoVJRyxXh5ztaQS4zeUqsUP0cDfjMLwoe4L4ZO32vU6GOp0c48k8EBw8tJaKhs

-- Dumped from database version 18.3 (Homebrew)
-- Dumped by pg_dump version 18.3 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: gerardozuniga
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';


ALTER DATABASE postgres OWNER TO gerardozuniga;

\unrestrict xwoVJRyxXh5ztaQS4zeUqsUP0cDfjMLwoe4L4ZO32vU6GOp0c48k8EBw8tJaKhs
\connect postgres
\restrict xwoVJRyxXh5ztaQS4zeUqsUP0cDfjMLwoe4L4ZO32vU6GOp0c48k8EBw8tJaKhs

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: gerardozuniga
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: car; Type: TABLE; Schema: public; Owner: postgres
--

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

--
-- Data for Name: car; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440000', 'Toyota', 'Corolla', 2020, 'Red', 'JTDBL40E799123456', 45000, 18999.99, false, '2026-04-15 17:56:31.26811-06', '2026-04-15 17:56:31.26811-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('f6d4daa6-d68f-487a-bc7f-ef31a8dd7747', 'Toyota', 'Rav4', 2022, 'White', 'JTDBL40E799123458', 15000, 18999.99, false, '0001-12-31 17:23:24-06:36:36 BC', '0001-12-31 17:23:24-06:36:36 BC');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('04c509aa-7038-4b6b-bf51-c6167118188a', 'Honda', 'Civic', 2023, 'Blue and Red', '2HGFE2F59PH123457', 8000, 21950.00, false, '0001-12-31 17:23:24-06:36:36 BC', '2026-04-15 19:01:50.566515-06');


--
-- Name: car car_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.car
    ADD CONSTRAINT car_pkey PRIMARY KEY (id);


--
-- Name: car car_vin_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.car
    ADD CONSTRAINT car_vin_key UNIQUE (vin);


--
-- PostgreSQL database dump complete
--

\unrestrict xwoVJRyxXh5ztaQS4zeUqsUP0cDfjMLwoe4L4ZO32vU6GOp0c48k8EBw8tJaKhs

