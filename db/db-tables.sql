-- Table: public.lists
-- DROP TABLE public.lists;
CREATE TABLE public.lists
(
    id integer NOT NULL DEFAULT nextval('lists_id_seq'::regclass),
    user_id integer NOT NULL,
    name character varying(256) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT lists_pkey PRIMARY KEY (id),
    CONSTRAINT lists_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
TABLESPACE pg_default;
ALTER TABLE public.lists OWNER to todo_admin;

-- Table: public.todos
-- DROP TABLE public.todos;
CREATE TABLE public.todos
(
    id integer NOT NULL DEFAULT nextval('todos_id_seq'::regclass),
    text character varying(256) COLLATE pg_catalog."default" NOT NULL,
    user_id integer NOT NULL,
    done boolean NOT NULL DEFAULT true,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT todos_pkey PRIMARY KEY (id),
    CONSTRAINT user_id FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
TABLESPACE pg_default;
ALTER TABLE public.todos OWNER to todo_admin;

-- Table: public.users
-- DROP TABLE public.users;
CREATE TABLE public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    username character varying(64) COLLATE pg_catalog."default" NOT NULL,
    salt character varying(16) COLLATE pg_catalog."default" NOT NULL,
    passhash character varying(60) COLLATE pg_catalog."default",
    deleted_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT username UNIQUE (username)
)

TABLESPACE pg_default;
ALTER TABLE public.users OWNER to todo_admin;
GRANT ALL ON TABLE public.users TO postgres;
GRANT ALL ON TABLE public.users TO todo_admin;
