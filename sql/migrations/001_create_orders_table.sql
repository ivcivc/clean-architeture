-- Migration: Create orders table
-- Created: 2024
-- Description: Creates the orders table with id, price, tax, and final_price columns

CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    final_price FLOAT NOT NULL,
    PRIMARY KEY (id)
);

-- Insert some example data for testing
INSERT IGNORE INTO orders (id, price, tax, final_price) VALUES 
('example-001', 100.0, 10.0, 110.0),
('example-002', 200.0, 20.0, 220.0),
('example-003', 300.0, 30.0, 330.0); 