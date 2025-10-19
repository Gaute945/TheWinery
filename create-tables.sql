DROP TABLE IF EXISTS albums;

CREATE TABLE
    albums (
        id INT AUTO_INCREMENT NOT NULL,
        title VARCHAR(128) NOT NULL,
        artist VARCHAR(255) NOT NULL,
        price DECIMAL(5, 2) NOT NULL,
        PRIMARY KEY (`id`)
    );

INSERT INTO
    albums (title, artist, price)
VALUES
    ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);

DROP TABLE IF EXISTS pages;

CREATE TABLE
    pages (
        title VARCHAR(128) NOT NULL,
        body VARCHAR(255) NOT NULL
    );

INSERT INTO
    pages (title, body)
VALUES
    ('page1', 'this is page 1!'),
    ('page2', 'this is page 2!'),
    ('page3', 'this is page 3!'),
    ('page4', 'this is page 4!'),
    ('page5', 'this is page 5!')