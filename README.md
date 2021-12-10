# Alert-System

An Alert notification service is an application which can receive alerts from certain alerting systems like System_X and System_Y and send these alerts to developers in the form of SMS and emails.

## External Componest used:

- **Postgres DB** (Data base)
- **Postman** (RestAPI testing)

## Pre-requisites:

```bash
foo@bar:~$ docker pull postgres
foo@bar:~$ docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```

### DB Set Up:

```sql
-- To create a database

-- Database: alertsystem

-- DROP DATABASE IF EXISTS alertsystem;

CREATE DATABASE alertsystem
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.developer

-- DROP TABLE IF EXISTS public.developer;

CREATE TABLE IF NOT EXISTS public.developer
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1000 MINVALUE 1000 MAXVALUE 10000000 CACHE 1 ),
    full_name character varying COLLATE pg_catalog."default",
    email character varying COLLATE pg_catalog."default",
    mobile text COLLATE pg_catalog."default",
    CONSTRAINT developer_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.developer
    OWNER to postgres;


-- Table: public.message

-- DROP TABLE IF EXISTS public.message;

CREATE TABLE IF NOT EXISTS public.message
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1000 MINVALUE 1000 MAXVALUE 10000000 CACHE 1 ),
    team_id bigint,
    content character varying COLLATE pg_catalog."default",
    title character varying COLLATE pg_catalog."default",
    msg_time timestamp without time zone,
    CONSTRAINT message_pkey PRIMARY KEY (id),
    CONSTRAINT team_id FOREIGN KEY (team_id)
        REFERENCES public.teams (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.message
    OWNER to postgres;


-- Table: public.teams

-- DROP TABLE IF EXISTS public.teams;

CREATE TABLE IF NOT EXISTS public.teams
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1000 MINVALUE 1000 MAXVALUE 99999999999999 CACHE 1 ),
    name character varying COLLATE pg_catalog."default",
    dept_name character varying COLLATE pg_catalog."default",
    developer_ids bigint[],
    CONSTRAINT teams_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.teams
    OWNER to postgres;

```

## Build and Run:

```bash
foo@bar:~$ git clone https://github.com/Aleena48/Alert-System.git
foo@bar:~$ cd Alert-System/cmd
foo@bar:~$ go build -o alertsystem
foo@bar:~$ ./alertsystem
```

## REST API Test:

Import the attached alert-system.postman_collection.json file in postman for workspace.
Logs are stored in logs folder for debugging.
