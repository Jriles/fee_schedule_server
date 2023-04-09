--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: attribute_values; Type: TABLE; Schema: public; Owner: jackriley
--

CREATE TABLE public.attribute_values (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title text NOT NULL,
    attribute_id uuid NOT NULL
);


ALTER TABLE public.attribute_values OWNER TO jackriley;

--
-- Name: attributes; Type: TABLE; Schema: public; Owner: jackriley
--

CREATE TABLE public.attributes (
    title text NOT NULL,
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    CONSTRAINT attributes_title_len CHECK ((length(title) < 200))
);


ALTER TABLE public.attributes OWNER TO jackriley;

--
-- Name: service_attribute_lines; Type: TABLE; Schema: public; Owner: jackriley
--

CREATE TABLE public.service_attribute_lines (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    service_id uuid NOT NULL,
    attribute_id uuid NOT NULL
);


ALTER TABLE public.service_attribute_lines OWNER TO jackriley;

--
-- Name: service_attribute_values; Type: TABLE; Schema: public; Owner: jackriley
--

CREATE TABLE public.service_attribute_values (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    line_id uuid NOT NULL,
    attribute_value_id uuid NOT NULL
);


ALTER TABLE public.service_attribute_values OWNER TO jackriley;

--
-- Name: service_variant_combination; Type: TABLE; Schema: public; Owner: jackriley
--

CREATE TABLE public.service_variant_combination (
    service_variant_id uuid NOT NULL,
    service_attribute_value_id uuid NOT NULL
);


ALTER TABLE public.service_variant_combination OWNER TO jackriley;

--
-- Name: service_variants; Type: TABLE; Schema: public; Owner: jackriley
--

CREATE TABLE public.service_variants (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    service_id uuid NOT NULL,
    state_cost integer NOT NULL,
    service_attribute_value_ids uuid[] NOT NULL,
    per_page_state_cost integer,
    country_code character varying(2) NOT NULL,
    currency_code character varying(3) NOT NULL
);


ALTER TABLE public.service_variants OWNER TO jackriley;

--
-- Name: services; Type: TABLE; Schema: public; Owner: jackriley
--

CREATE TABLE public.services (
    title text NOT NULL,
    id uuid DEFAULT gen_random_uuid() NOT NULL
);


ALTER TABLE public.services OWNER TO jackriley;

--
-- Data for Name: attribute_values; Type: TABLE DATA; Schema: public; Owner: jackriley
--

COPY public.attribute_values (id, title, attribute_id) FROM stdin;
e4ae3ed0-c89d-45ae-a5c9-1e3c166f1ea0	Alabama	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
47098832-9e60-445c-90a0-0f0b1d8c25fc	Delaware	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
5865863a-0803-48f3-a9b5-e65a89b5fc25	New Jersey	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
dfb12ac0-d7e3-4834-83e1-a4483c370cef	Texas	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
644aa685-af46-4f2e-8ecd-36d864070af7	Wisconsin	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
3e93aac9-5f37-4e78-8158-11613d9cc9c6	New York	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
660016e3-8c6f-4a05-b7de-733ae28be0ac	S Corp	98b1e939-fae7-45aa-b7cc-73c7f5c8e40e
2e4999f6-3e85-439d-b792-92802dcd1472	LP	98b1e939-fae7-45aa-b7cc-73c7f5c8e40e
441b2154-2409-49c1-bef1-484a82a33631	PartnerShip	98b1e939-fae7-45aa-b7cc-73c7f5c8e40e
28837049-c805-4d8e-8506-6807083c4226	North Dakota	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
a608b7e1-2067-42e2-a5e1-749be5f152f0	1 Day	b2ee5840-5c32-4703-a18d-abd1e257fe70
87af7357-6ddc-4a17-aa50-1d38e719b1d1	10000	0ec3827f-4c72-4b73-b1b8-c766ee9716a1
dd6ca79e-f990-49e7-bbbb-b7cff19637cd	Not-a-corp	98b1e939-fae7-45aa-b7cc-73c7f5c8e40e
74a66d4e-5569-4085-a2aa-7539d6b2ad09	Expedited	b2ee5840-5c32-4703-a18d-abd1e257fe70
\.


--
-- Data for Name: attributes; Type: TABLE DATA; Schema: public; Owner: jackriley
--

COPY public.attributes (title, id) FROM stdin;
Entity Type	98b1e939-fae7-45aa-b7cc-73c7f5c8e40e
Filing Speed	b2ee5840-5c32-4703-a18d-abd1e257fe70
Stock Count	0ec3827f-4c72-4b73-b1b8-c766ee9716a1
Jurisdiction	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
\.


--
-- Data for Name: service_attribute_lines; Type: TABLE DATA; Schema: public; Owner: jackriley
--

COPY public.service_attribute_lines (id, service_id, attribute_id) FROM stdin;
b7daab28-0296-4c86-9639-2b7a92530091	418d4bdc-1115-4625-a7fe-cd22b10755fe	b2ee5840-5c32-4703-a18d-abd1e257fe70
3ee73c25-b1fe-4d59-8190-ad7d93aca69c	c726dcb1-41a3-4c14-9cdd-705f19e5203d	0ec3827f-4c72-4b73-b1b8-c766ee9716a1
d10eb74a-d18c-4f33-980a-58d9d1e85b9a	c726dcb1-41a3-4c14-9cdd-705f19e5203d	2d235cd6-3a77-4015-8aae-7cb8e2e91c64
\.


--
-- Data for Name: service_attribute_values; Type: TABLE DATA; Schema: public; Owner: jackriley
--

COPY public.service_attribute_values (id, line_id, attribute_value_id) FROM stdin;
9ced79f5-8cc6-4fc2-b8c3-050a35e1ee04	3ee73c25-b1fe-4d59-8190-ad7d93aca69c	87af7357-6ddc-4a17-aa50-1d38e719b1d1
b58cb1e7-21d6-436b-8516-efae544fe201	b7daab28-0296-4c86-9639-2b7a92530091	a608b7e1-2067-42e2-a5e1-749be5f152f0
962ab42f-359e-4bab-918f-787e41250853	b7daab28-0296-4c86-9639-2b7a92530091	74a66d4e-5569-4085-a2aa-7539d6b2ad09
\.


--
-- Data for Name: service_variant_combination; Type: TABLE DATA; Schema: public; Owner: jackriley
--

COPY public.service_variant_combination (service_variant_id, service_attribute_value_id) FROM stdin;
c30b4e5a-0a9a-4150-a9dc-949258e7b1f3	b58cb1e7-21d6-436b-8516-efae544fe201
\.


--
-- Data for Name: service_variants; Type: TABLE DATA; Schema: public; Owner: jackriley
--

COPY public.service_variants (id, service_id, state_cost, service_attribute_value_ids, per_page_state_cost, country_code, currency_code) FROM stdin;
c30b4e5a-0a9a-4150-a9dc-949258e7b1f3	418d4bdc-1115-4625-a7fe-cd22b10755fe	15000054	{b58cb1e7-21d6-436b-8516-efae544fe201}	0	US	USD
\.


--
-- Data for Name: services; Type: TABLE DATA; Schema: public; Owner: jackriley
--

COPY public.services (title, id) FROM stdin;
Formation	418d4bdc-1115-4625-a7fe-cd22b10755fe
Amendment foreign	c726dcb1-41a3-4c14-9cdd-705f19e5203d
\.


--
-- Name: attribute_values attribute_values_pkey; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.attribute_values
    ADD CONSTRAINT attribute_values_pkey PRIMARY KEY (id);


--
-- Name: attribute_values attribute_values_title_len; Type: CHECK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE public.attribute_values
    ADD CONSTRAINT attribute_values_title_len CHECK ((length(title) < 200)) NOT VALID;


--
-- Name: attributes attributes_pkey; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.attributes
    ADD CONSTRAINT attributes_pkey PRIMARY KEY (id);


--
-- Name: service_variant_combination service attr val, service variant, combo unique; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_variant_combination
    ADD CONSTRAINT "service attr val, service variant, combo unique" UNIQUE (service_variant_id, service_attribute_value_id);


--
-- Name: service_attribute_lines service_attribute_line_pkey; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_lines
    ADD CONSTRAINT service_attribute_line_pkey PRIMARY KEY (id);


--
-- Name: service_attribute_values service_attribute_values_pkey; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_values
    ADD CONSTRAINT service_attribute_values_pkey PRIMARY KEY (id);


--
-- Name: service_variants service_variants_pkey; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_variants
    ADD CONSTRAINT service_variants_pkey PRIMARY KEY (id);


--
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- Name: services services_title_len; Type: CHECK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE public.services
    ADD CONSTRAINT services_title_len CHECK ((length(title) < 200)) NOT VALID;


--
-- Name: service_attribute_lines unique_attribute_line_per_service; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_lines
    ADD CONSTRAINT unique_attribute_line_per_service UNIQUE (service_id, attribute_id) INCLUDE (service_id, attribute_id);


--
-- Name: attribute_values unique_attribute_values; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.attribute_values
    ADD CONSTRAINT unique_attribute_values UNIQUE (title, attribute_id) INCLUDE (title, attribute_id);


--
-- Name: attributes unique_attributes; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.attributes
    ADD CONSTRAINT unique_attributes UNIQUE (title) INCLUDE (title);


--
-- Name: service_attribute_values unique_service_attribute_values_to_line; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_values
    ADD CONSTRAINT unique_service_attribute_values_to_line UNIQUE (line_id, attribute_value_id) INCLUDE (line_id, attribute_value_id);


--
-- Name: service_variants unique_service_variants; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_variants
    ADD CONSTRAINT unique_service_variants UNIQUE (service_id, service_attribute_value_ids) INCLUDE (service_id, service_attribute_value_ids);


--
-- Name: services unique_services; Type: CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT unique_services UNIQUE (title) INCLUDE (title);


--
-- Name: fki_service_id_fkey; Type: INDEX; Schema: public; Owner: jackriley
--

CREATE INDEX fki_service_id_fkey ON public.service_variants USING btree (service_id);


--
-- Name: fki_service_line_id; Type: INDEX; Schema: public; Owner: jackriley
--

CREATE INDEX fki_service_line_id ON public.service_attribute_values USING btree (line_id);


--
-- Name: service_attribute_lines attribute_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_lines
    ADD CONSTRAINT attribute_id_fkey FOREIGN KEY (attribute_id) REFERENCES public.attributes(id) ON DELETE CASCADE NOT VALID;


--
-- Name: attribute_values attribute_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.attribute_values
    ADD CONSTRAINT attribute_id_fkey FOREIGN KEY (attribute_id) REFERENCES public.attributes(id) ON DELETE CASCADE NOT VALID;


--
-- Name: service_attribute_values attribute_value_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_values
    ADD CONSTRAINT attribute_value_id_fkey FOREIGN KEY (attribute_value_id) REFERENCES public.attribute_values(id) ON DELETE CASCADE NOT VALID;


--
-- Name: service_attribute_values line_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_values
    ADD CONSTRAINT line_id_fkey FOREIGN KEY (line_id) REFERENCES public.service_attribute_lines(id) NOT VALID;


--
-- Name: service_variants service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_variants
    ADD CONSTRAINT service_id_fkey FOREIGN KEY (service_id) REFERENCES public.services(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- Name: service_attribute_lines service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_attribute_lines
    ADD CONSTRAINT service_id_fkey FOREIGN KEY (service_id) REFERENCES public.services(id) ON DELETE CASCADE NOT VALID;


--
-- Name: service_variant_combination service_variant_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jackriley
--

ALTER TABLE ONLY public.service_variant_combination
    ADD CONSTRAINT service_variant_fkey FOREIGN KEY (service_variant_id) REFERENCES public.service_variants(id) ON DELETE CASCADE NOT VALID;


--
-- PostgreSQL database dump complete
--

