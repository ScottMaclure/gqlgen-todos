-- SEQUENCE: public.lists_id_seq

-- DROP SEQUENCE public.lists_id_seq;

CREATE SEQUENCE public.lists_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.lists_id_seq
    OWNER TO todo_admin;

-- SEQUENCE: public.todos_id_seq

-- DROP SEQUENCE public.todos_id_seq;

CREATE SEQUENCE public.todos_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.todos_id_seq
    OWNER TO todo_admin;

-- SEQUENCE: public.users_id_seq

-- DROP SEQUENCE public.users_id_seq;

CREATE SEQUENCE public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

ALTER SEQUENCE public.users_id_seq
    OWNER TO todo_admin;        