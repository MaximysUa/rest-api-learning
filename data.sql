CREATE TABLE public.author(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

--Many-to-Many
CREATE TABLE public.book_authors(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID NOT NULL ,
    author_id UUID NOT NULL ,

    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id) ,
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id),
    CONSTRAINT book_author_unique UNIQUE (book_id, author_id)

);



INSERT INTO author(name) VALUES ('Народ');
INSERT INTO author(name) VALUES ('Джолиан Роулинг');
INSERT INTO author(name) VALUES ('Джек Лондон');

INSERT INTO book(name) VALUES ('колобок');
INSERT INTO book(name) VALUES ('Гарри поттер');
INSERT INTO book(name) VALUES ('Бриллианты');

INSERT INTO book_authors(book_id, author_id) VALUES ('16090d67-6f7d-4156-bd9d-238d9d14c964', 'f2c6a54d-e3be-4028-a561-8f035ee5f9dd');
INSERT INTO book_authors(book_id, author_id) VALUES ('16090d67-6f7d-4156-bd9d-238d9d14c964', '2ee2a987-9a2b-4c50-ad32-890c14a2ec59');
INSERT INTO book_authors(book_id, author_id) VALUES ('16090d67-6f7d-4156-bd9d-238d9d14c964', '5aaaa4a6-e8df-44fc-8e88-896c3880933c');
INSERT INTO book_authors(book_id, author_id) VALUES ('fb72f699-a1ee-4bf8-b1cc-664dcb624f8b', 'f2c6a54d-e3be-4028-a561-8f035ee5f9dd');
INSERT INTO book_authors(book_id, author_id) VALUES ('fb72f699-a1ee-4bf8-b1cc-664dcb624f8b', '2ee2a987-9a2b-4c50-ad32-890c14a2ec59');