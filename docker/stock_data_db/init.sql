CREATE DATABASE IF NOT EXISTS stock_data_db;
USE stock_data_db;

CREATE USER IF NOT EXISTS 'finet_app'@'%' IDENTIFIED BY 'finet_password';
GRANT ALL PRIVILEGES ON stock_data_db.* TO 'finet_app'@'%';
FLUSH PRIVILEGES;

CREATE TABLE IF NOT EXISTS stock_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ticker VARCHAR(10) NOT NULL,
    date DATE NOT NULL,
    open DOUBLE,
    high DOUBLE,
    low DOUBLE,
    close DOUBLE,
    adj_close DOUBLE,
    volume BIGINT,
    dividend DOUBLE,
    
    UNIQUE (ticker, date)
);
