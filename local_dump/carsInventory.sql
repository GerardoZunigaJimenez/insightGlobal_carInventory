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
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440001', 'Toyota', 'Corolla', 2020, 'Red', 'JTDBL40E799123457', 45000, 18999.99, false, '2026-04-12 09:14:22.11811-06', '2026-04-12 09:14:22.11811-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440002', 'Honda', 'Civic', 2021, 'Blue', '2HGFC2F69MH123458', 32000, 21450.00, false, '2026-04-11 14:03:51.42133-06', '2026-04-11 14:03:51.42133-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440003', 'Ford', 'Focus', 2019, 'Black', '1FADP3F25KL123459', 58000, 15990.50, false, '2026-04-10 18:47:09.77645-06', '2026-04-10 18:47:09.77645-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440004', 'Chevrolet', 'Malibu', 2022, 'White', '1G1ZD5ST2NF123450', 18000, 23999.00, false, '2026-04-09 07:22:13.55421-06', '2026-04-09 07:22:13.55421-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440005', 'Nissan', 'Sentra', 2020, 'Gray', '3N1AB7AP8KY123451', 41000, 17800.75, false, '2026-04-08 11:58:40.93412-06', '2026-04-08 11:58:40.93412-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440006', 'Mazda', '3', 2023, 'Silver', 'JM1BPBEM0P123452', 12000, 26850.00, false, '2026-04-07 16:31:05.26789-06', '2026-04-07 16:31:05.26789-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440007', 'Hyundai', 'Elantra', 2021, 'Red', 'KMHLM4AG2MU123453', 29000, 20500.00, false, '2026-04-06 10:12:44.99102-06', '2026-04-06 10:12:44.99102-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440008', 'Kia', 'Forte', 2022, 'Blue', '3KPF24AD6NE123454', 22000, 19875.30, false, '2026-04-05 13:44:28.10566-06', '2026-04-05 13:44:28.10566-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440009', 'Subaru', 'Impreza', 2020, 'Green', 'JF1GPAN67LH123455', 36000, 21999.99, false, '2026-04-04 19:05:16.33210-06', '2026-04-04 19:05:16.33210-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440010', 'Volkswagen', 'Jetta', 2021, 'White', '3VWCB7BU5MM123456', 27000, 22750.00, false, '2026-04-03 08:20:57.66444-06', '2026-04-03 08:20:57.66444-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440011', 'BMW', '320i', 2023, 'Black', 'WBA53AK02P123457', 9000, 38999.99, false, '2026-04-02 15:39:11.00988-06', '2026-04-02 15:39:11.00988-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440012', 'Audi', 'A4', 2022, 'Silver', 'WAUENAF48NA123458', 14000, 42950.00, false, '2026-04-01 12:10:33.48765-06', '2026-04-01 12:10:33.48765-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440013', 'Mercedes-Benz', 'C300', 2021, 'Gray', 'W1KAF4HB1MR123459', 25000, 45999.95, false, '2026-03-31 17:25:48.77123-06', '2026-03-31 17:25:48.77123-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440014', 'Tesla', 'Model 3', 2023, 'White', '5YJ3E1EA7PF123460', 7000, 39999.00, false, '2026-03-30 09:48:02.21459-06', '2026-03-30 09:48:02.21459-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440015', 'Toyota', 'Camry', 2022, 'Blue', '4T1B11HK5NU123461', 16000, 27999.50, false, '2026-03-29 14:16:19.99010-06', '2026-03-29 14:16:19.99010-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440016', 'Honda', 'Accord', 2020, 'Black', '1HGCV1F14LA123462', 39000, 24900.00, false, '2026-03-28 11:55:03.61877-06', '2026-03-28 11:55:03.61877-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440017', 'Ford', 'Escape', 2021, 'Red', '1FMCU0F61MU123463', 30000, 26800.25, false, '2026-03-27 18:07:45.13056-06', '2026-03-27 18:07:45.13056-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440018', 'Chevrolet', 'Equinox', 2023, 'Green', '2GNAXXEV4P6123464', 11000, 31500.00, false, '2026-03-26 07:33:21.44512-06', '2026-03-26 07:33:21.44512-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440019', 'Nissan', 'Altima', 2019, 'Silver', '1N4BL4BV3KC123465', 52000, 17150.00, false, '2026-03-25 13:02:58.88331-06', '2026-03-25 13:02:58.88331-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440020', 'Mazda', 'CX-5', 2022, 'Pearl White', 'JM3KFBDM7N123466', 21000, 32999.99, false, '2026-03-24 20:45:10.77201-06', '2026-03-24 20:45:10.77201-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440021', 'Kia', 'Sportage', 2021, 'Orange', 'KNDPMCAC7M123467', 26000, 28650.00, false, '2026-03-23 09:17:34.55118-06', '2026-03-23 09:17:34.55118-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440022', 'Subaru', 'Outback', 2020, 'Brown', '4S4BTALC4L123468', 47000, 28999.00, false, '2026-03-22 16:28:49.00664-06', '2026-03-22 16:28:49.00664-06');
INSERT INTO public.car (id, make, model, year, color, vin, mileage, price, disabled, created_at, updated_at) VALUES ('550e8400-e29b-41d4-a716-446655440023', 'Volkswagen', 'Passat', 2023, 'Dark Blue', '1VWAB7A32PC123469', 6000, 34999.95, false, '2026-03-21 12:59:05.33770-06', '2026-03-21 12:59:05.33770-06');