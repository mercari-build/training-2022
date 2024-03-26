CREATE TABLE category (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255)
);
CREATE TABLE items (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255),
    category_id INTEGER,
    image_name VARCHAR(255),
    FOREIGN KEY (category_id) REFERENCES category(id)
);
