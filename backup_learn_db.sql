--
-- PostgreSQL database dump
--

-- Dumped from database version 17.6
-- Dumped by pg_dump version 17.4

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: rentals; Type: TABLE; Schema: public; Owner: m16yusuf
--

CREATE TABLE public.rentals (
    id integer NOT NULL,
    image text,
    rentals_name character varying(50) NOT NULL,
    user_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone
);


ALTER TABLE public.rentals OWNER TO m16yusuf;

--
-- Name: rentals_id_seq; Type: SEQUENCE; Schema: public; Owner: m16yusuf
--

ALTER TABLE public.rentals ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.rentals_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: userse; Type: TABLE; Schema: public; Owner: m16yusuf
--

CREATE TABLE public.userse (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at timestamp without time zone
);


ALTER TABLE public.userse OWNER TO m16yusuf;

--
-- Name: userse_id_seq; Type: SEQUENCE; Schema: public; Owner: m16yusuf
--

CREATE SEQUENCE public.userse_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.userse_id_seq OWNER TO m16yusuf;

--
-- Name: userse_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: m16yusuf
--

ALTER SEQUENCE public.userse_id_seq OWNED BY public.userse.id;


--
-- Name: userse id; Type: DEFAULT; Schema: public; Owner: m16yusuf
--

ALTER TABLE ONLY public.userse ALTER COLUMN id SET DEFAULT nextval('public.userse_id_seq'::regclass);


--
-- Data for Name: rentals; Type: TABLE DATA; Schema: public; Owner: m16yusuf
--

COPY public.rentals (id, image, rentals_name, user_id, created_at, updated_at) FROM stdin;
2	mobil.png	Honda	1	2025-08-26 16:41:09.434586	\N
3	motor.jpeg	supra	2	2025-08-26 16:41:09.434586	\N
4	gambar.jpeg	ferrari	1	2025-08-26 16:47:31.467983	\N
7	Ferrari_2019.png	Ferrari tahun 2019	1	2025-09-03 10:00:17.82424	\N
8	supra-2018.png	Rental Bandung supra	2	2025-09-03 22:30:25.324667	2025-09-03 22:41:05.126268
10	mazda.png	Rntal sugiono	2	2025-09-04 02:12:58.772575	\N
11	mazda-2015.png	Rental sugiono	1	2025-09-04 02:13:22.730373	\N
9	mazda-2015.png	Rental sugiono bin miharza	1	2025-09-04 02:05:16.213901	2025-09-04 02:14:42.452271
\.


--
-- Data for Name: userse; Type: TABLE DATA; Schema: public; Owner: m16yusuf
--

COPY public.userse (id, username, password, created_at, update_at) FROM stdin;
1	andi	827ccb0eea8a706c4c34a16891f84e7b	2025-08-26 16:34:36.716038	\N
2	bonai	827ccb0eea8a706c4c34a16891f84e7b	2025-08-26 16:34:36.716038	\N
\.


--
-- Name: rentals_id_seq; Type: SEQUENCE SET; Schema: public; Owner: m16yusuf
--

SELECT pg_catalog.setval('public.rentals_id_seq', 11, true);


--
-- Name: userse_id_seq; Type: SEQUENCE SET; Schema: public; Owner: m16yusuf
--

SELECT pg_catalog.setval('public.userse_id_seq', 2, true);


--
-- Name: rentals rentals_name_unique_constraint; Type: CONSTRAINT; Schema: public; Owner: m16yusuf
--

ALTER TABLE ONLY public.rentals
    ADD CONSTRAINT rentals_name_unique_constraint UNIQUE (rentals_name);


--
-- Name: userse userse_pkey; Type: CONSTRAINT; Schema: public; Owner: m16yusuf
--

ALTER TABLE ONLY public.userse
    ADD CONSTRAINT userse_pkey PRIMARY KEY (id);


--
-- Name: userse userse_username_key; Type: CONSTRAINT; Schema: public; Owner: m16yusuf
--

ALTER TABLE ONLY public.userse
    ADD CONSTRAINT userse_username_key UNIQUE (username);


--
-- Name: rentals rentals_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: m16yusuf
--

ALTER TABLE ONLY public.rentals
    ADD CONSTRAINT rentals_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.userse(id);


--
-- PostgreSQL database dump complete
--

