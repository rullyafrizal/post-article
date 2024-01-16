CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200),
    content TEXT,
    category VARCHAR(100),
    created_date TIMESTAMP,
    updated_date TIMESTAMP,
    status VARCHAR(100) CHECK (status IN ('Publish', 'Draft', 'Thrash'))
);