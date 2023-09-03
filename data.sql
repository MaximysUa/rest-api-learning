CREATE TABLE public.author(
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book(
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            name VARCHAR(100) NOT NULL ,
                            author_id UUID NOT NULL ,
                            CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author(name) VALUES ('Народ');
INSERT INTO author(name) VALUES ('Джолиан Роулинг');
INSERT INTO author(name) VALUES ('Джек Лондон');

INSERT INTO book(name, author_id) VALUES ('колобок', 'd34e541a-4166-442b-9e87-a1f26080ab0f');
INSERT INTO book(name, author_id) VALUES ('Гарри поттер', 'd4f8ad56-ff44-4c4f-b8ed-e4ee1dc27acd');
INSERT INTO book(name, author_id) VALUES ('Бриллианты', '810e18f9-0b5c-42a4-8d97-e4e673c97e3a');