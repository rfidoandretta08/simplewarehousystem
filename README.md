database hanya perlu diinput dummy data dengan raw manual mysql untuk bisa mencoba fitur pada postman GET dan PUT 
dummy data disediakan sbb : 
USE day23;
SHOW TABLES;

INSERT INTO products (name, description, price, category) VALUES
('Laptop ASUS', 'Laptop dengan prosesor Intel i7 dan RAM 16GB', 12000000.00, 'Elektronik'),
('Kulkas Samsung', 'Kulkas dua pintu dengan teknologi Inverter', 5000000.00, 'Peralatan Rumah Tangga'),
('Sepatu Nike', 'Sepatu olahraga dengan teknologi Flyknit', 800000.00, 'Olahraga'),
('Smartphone Xiaomi', 'Smartphone dengan layar AMOLED 6.5 inci', 4000000.00, 'Elektronik');

INSERT INTO inventories (product_id, quantity, location) VALUES
(1, 50, 'Gudang A'),
(2, 30, 'Gudang B'),
(3, 100, 'Gudang A'),
(4, 70, 'Gudang C');

INSERT INTO orders (product_id, quantity, order_date) VALUES
(1, 2, '2025-02-01'),
(3, 5, '2025-02-02'),
(2, 1, '2025-02-03'),
(4, 3, '2025-02-04');


SELECT * FROM products;
SELECT * FROM inventories;
SELECT * FROM orders;
SELECT * FROM orders WHERE product_id = 1;
SHOW TABLES;

untuk database ini bisa membuat sendiri dengan mengganti pada
config database.go dsn := "root:pass my sql anda@tcp(127.0.0.1:3306)/database anda?charset=utf8mb4&parseTime=True&loc=Local"

untuk data ini adalah data one to many jadi 1 produk id bisa memiliki relasi ke lebih dari 1 table 
1 produk id bisa memiliki lebih dari 1 order dan lebih dari 1 tempat gudang 

