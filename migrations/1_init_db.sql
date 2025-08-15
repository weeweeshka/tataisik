CREATE TABLE IF NOT EXISTS films
(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    year_of_prod INTEGER NOT NULL CHECK(year_of_prod >= 1900),
    imdb FLOAT NOT NULL CHECK(imdb > 0 and imdb <= 10),
    description TEXT NOT NULL,
    country TEXT[] NOT NULL,
    genre TEXT[] NOT NULL,
    film_director TEXT NOT NULL,
    screenwriter TEXT NOT NULL,
    budget INTEGER NOT NULL,
    collection INTEGER NOT NULL
);