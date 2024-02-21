CREATE TABLE book (
    id uuid PRIMARY KEY NOT NULL,
    book_name varchar(30),
    author_name varchar(30) NOT NULL, 
    page_number int NOT NULL
);


