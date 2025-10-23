DROP TABLE IF EXISTS wines;

CREATE TABLE wines (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    grape VARCHAR(128),
    origin VARCHAR(256),
    producer VARCHAR(128),
    vintage INT(4),
    taste TEXT,
    color TEXT,
    smell TEXT,
    acidity DECIMAL(4,2),
    sweetness DECIMAL(6,2),
    price DECIMAL(7,2),
    iSwinescale DECIMAL(3,1) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO wines (
    title, grape, origin, producer, vintage,
    taste, color, smell, acidity, sweetness, price, iSwinescale
) VALUES
    (
        'Clos de la Coulée de Serrant',
        'Chenin Blanc',
        'Loire Valley, France',
        'Nicolas Joly',
        2019,
        'Rich texture with layers of honey, almond, and ripe pear.',
        'Deep golden hue',
        'Floral, quince, light oxidation',
        6.8,
        4.2,
        89.90,
        9.3
    ),
    (
        'Barolo Cannubi',
        'Nebbiolo',
        'Piedmont, Italy',
        'Paolo Scavino',
        2016,
        'Firm tannins, red cherry, tar, and rose petals.',
        'Garnet red',
        'Violets, truffle, dried herbs',
        5.5,
        1.2,
        74.50,
        9.0
    ),
    (
        'Riesling Kabinett',
        'Riesling',
        'Mosel, Germany',
        'Dr. Loosen',
        2021,
        'Lively acidity, green apple, light sweetness.',
        'Pale straw',
        'Lime zest, slate, green apple',
        8.9,
        35.0,
        18.90,
        8.7
    ),
    (
        'Malbec Reserva',
        'Malbec',
        'Mendoza, Argentina',
        'Catena Zapata',
        2020,
        'Full-bodied, ripe plum, cocoa, and vanilla oak.',
        'Deep violet',
        'Blackberry, chocolate, spice',
        5.1,
        3.5,
        24.50,
        8.2
    ),
    (
        'Sauternes Château d’Yquem',
        'Sémillon / Sauvignon Blanc',
        'Bordeaux, France',
        'Château d’Yquem',
        2011,
        'Intense sweetness balanced by acidity, apricot, and honey.',
        'Amber gold',
        'Apricot jam, saffron, honeycomb',
        6.2,
        140.0,
        420.00,
        9.8
    );
