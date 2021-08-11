CREATE DATABASE ALMACENES;

USE ALMACENES;

CREATE TABLE Cajeros(
	Cajero INT NOT NULL AUTO_INCREMENT,
	NomApels VARCHAR(255),
	PRIMARY KEY (Cajero)
);

CREATE TABLE Productos(
	Producto INT NOT NULL AUTO_INCREMENT,
	Nombre VARCHAR(255),
	Precio FLOAT(10,2),
	PRIMARY KEY (Producto)
);

CREATE TABLE Maquinas_Registradoras(
	Maquina INT NOT NULL AUTO_INCREMENT,
	Piso INT,
	PRIMARY KEY (Maquina)
);

CREATE TABLE Venta(
	Cajero INT NOT NULL,
	Maquina INT NOT NULL,
	Producto INT NOT NULL,
	FOREIGN KEY (Cajero) REFERENCES Cajeros(Cajero),
	FOREIGN KEY (Maquina) REFERENCES Maquinas_Registradoras(Maquina),
	FOREIGN KEY (Producto) REFERENCES Productos(Producto)
);

--QUERIES

-- (a) poblar todas las tablas.

--INSERTA VALORES EN LA TABLA DE CAJEROS
INSERT INTO Cajeros (NomApels) VALUES 
	("Andres"), 
	("Jorge"), 
	("Tomas"), 
	("Ivan");

--INSERTA VALORES EN LA TABLA DE PRODUCTOS
INSERT INTO Productos 
	(Nombre, Precio)
VALUES 
	("hoja", 1.50), 
	("camiseta", 10.50), 
	("pantalon", 15.95), 
	("calcetin", 3.34),
	("computadora", 10999.99),
	("lapiz", 4.50),
	("cobija", 100.55);
	
--INSERTA VALORES EN LA TABLA DE MAQUINAS
INSERT INTO Maquinas_Registradoras (Piso) VALUES (1), (1), (2), (2);

--INSERTA VALORES EN LA TABLA DE VENTA
INSERT INTO Venta 
	(Cajero, Maquina, Producto)
VALUES
	(3,2,5),
	(3,2,1),
	(1,3,2),
	(1,4,1),
	(2,1,4),
	(4,3,4),
	(2,1,7),
	(1,4,3),
	(2,1,1),
	(1,4,7),
	(3,2,3),
	(4,2,4),
	(1,3,6),
	(2,1,6),
	(3,2,7);



-- (b) Número de ventas de cada producto
SELECT p.Nombre AS 'Producto', COUNT(*) AS 'Vendido'
FROM Venta v
INNER JOIN Productos p ON p.Producto = v.Producto
GROUP BY v.Producto
ORDER BY Vendido DESC;

-- (c) informe completo de ventas
SELECT  c.NomApels AS 'Cajero', 
	p.Nombre AS 'Producto', 
	p.Precio AS 'Precio producto', 
	m.Piso AS 'Piso registradora'
FROM Venta v
INNER JOIN Cajeros c ON v.Cajero = c.Cajero
INNER JOIN Productos p ON v.Producto = p.Producto
INNER JOIN Maquinas_Registradoras m ON v.Maquina = m.Maquina;

-- (d) ventas totales realizadas en cada piso
SELECT m.Piso,
	COUNT(*) AS 'Ventas'
FROM Venta v
INNER JOIN Maquinas_Registradoras m ON v.Maquina = m.Maquina
GROUP BY m.Piso;

-- (e) código y nombre de cada cajero junto con el importe total de sus ventas
SELECT c.Cajero AS 'Codigo cajero' , 
	c.NomApels AS 'Cajero',
	SUM(p.Precio) AS 'Vendio'
FROM Venta v
INNER JOIN Cajeros c ON v.Cajero = c.Cajero
INNER JOIN Productos p ON v.Producto = p.Producto
GROUP BY c.Cajero;

-- (f) código y nombre de aquellos cajeros que hayan realizado ventas en pisos cuyas ventas totales sean inferiores a los 5000 pesos.
SELECT C.Cajero AS 'Codigo cajero', C.NomApels AS 'Cajero'
FROM Venta V
INNER JOIN Maquinas_Registradoras MR ON V.Maquina = MR.Maquina
INNER JOIN Cajeros C ON V.Cajero = C.Cajero
INNER JOIN (
	SELECT m.Piso, SUM(p.Precio) AS 'Vendido'
	FROM Venta v
	INNER JOIN Cajeros c ON v.Cajero = c.Cajero
	INNER JOIN Productos p ON v.Producto = p.Producto
	INNER JOIN Maquinas_Registradoras m ON v.Maquina = m.Maquina
	GROUP BY m.Piso
) AS PisoVenta
ON PisoVenta.Piso = MR.Piso
WHERE PisoVenta.Vendido < 5000
GROUP BY C.NomApels, C.Cajero;









