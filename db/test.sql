-- create user
INSERT INTO public.users(
	username, 
	salt, 
	passhash
) VALUES (
	'hans_gruber', 
	'GlrkUDBEmlyVMBth', 
	'$2a$14$naxeroSLyRgojKFw3EBMoe8vMXcHLyJRXVNMzP0XLwi8SwEO4jNZ.'
);

-- update user
UPDATE public.users SET
created_at = '2020-09-04 13:16:45.119229+08'::timestamp with time zone WHERE
id = 2;

-- create todo items
INSERT INTO public.todos(text, user_id) VALUES ('Integrate with GraphQL server', 1);
INSERT INTO public.todos(text, user_id) VALUES ('My TODO Item A', 1);
INSERT INTO public.todos(text, user_id) VALUES ('My TODO Item B', 1);
INSERT INTO public.todos(text, user_id) VALUES ('My TODO Item C', 1);