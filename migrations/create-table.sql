DROP TABLE IF EXISTS games;
CREATE TABLE games (
    game_id INT PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    releaseDate DATE NOT NULL,
    price DECIMAL(5, 2) NOT NULL
);

CREATE TABLE reviews (
    review_id INT PRIMARY KEY,
    game_id INT NOT NULL,
    rating INT NOT NULL,
    description TEXT,
    created_at TIME NOT NULL,
    updated_at TIME,
    CONSTRAINT fk_game_id
        FOREIGN KEY(game_id)
        REFERENCES games(game_id)
);

INSERT INTO reviews 
    (review_id, game_id, rating, description, created_at, updated_at)
VALUES
    (0, 1, 9, 'Bloodborne Review', '18:11:00', '18:11:00');

INSERT INTO games 
    (id, title, genre, releaseDate, price) 
VALUES
    ('Elden Ring', 'Action RPG', '2022-02-25', 59.99),
    ('Bloodborne', 'Action RPG', '2015-03-24', 29.99),
    ('God of War', 'Action-Adventure', '2018-04-20', 59.99);