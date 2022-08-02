-- --------------------------------------------------------
-- Hôte:                         127.0.0.1
-- Version du serveur:           PostgreSQL 12.11 on x86_64-pc-linux-musl, compiled by gcc (Alpine 11.2.1_git20220219) 11.2.1 20220219, 64-bit
-- SE du serveur:                
-- HeidiSQL Version:             12.0.0.6536
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Listage des données de la table public.cars : 10 rows
/*!40000 ALTER TABLE "cars" DISABLE KEYS */;
INSERT INTO "cars" ("id", "matriculation_number", "brand", "model", "year") VALUES
	(1, 'AB123CD', 'Peugeot', '405', 1999),
	(2, 'EF456GH', 'Renault', 'Megane', 1996),
	(3, 'IJ789KL', 'Toyota', 'Yaris', 2001),
	(4, 'MN147OP', 'Volkswagen', 'Polo', 1997),
	(5, 'HI987JK', 'Fiat', '500', 2008),
	(6, 'QR258ST', 'Peugeot', '406', 1998),
	(7, 'UV369WX', 'Renault', 'Megane', 1997),
	(8, 'ZA321BC', 'Toyota', 'Yaris', 2002),
	(9, 'DE654FG', 'Volkswagen', 'Polo', 1998),
	(10, 'LM753OP', 'Fiat', '500', 2008);
/*!40000 ALTER TABLE "cars" ENABLE KEYS */;

-- Listage des données de la table public.customer2car : 5 rows
/*!40000 ALTER TABLE "customer2car" DISABLE KEYS */;
INSERT INTO "customer2car" ("id", "customer_id", "car_id") VALUES
	(3, 2, 5),
	(4, 2, 6),
	(5, 1, 8),
	(6, 1, 9),
	(7, 1, 10);
/*!40000 ALTER TABLE "customer2car" ENABLE KEYS */;

-- Listage des données de la table public.customers : 7 rows
/*!40000 ALTER TABLE "customers" DISABLE KEYS */;
INSERT INTO "customers" ("id", "email", "name", "firstname", "birthday", "driving_licence_number") VALUES
	(1, 'jean.dupont@domain.com', 'Dupont', 'Jean', '1990-12-13', 'JEAND657054SM9IJ'),
	(2, 'olivier.duchene@a-pro.fr', 'Duchene', 'Olivier', '1985-06-20', 'OLIVD657055SM9IJ'),
	(4, 'test1659436392@a-pro.fr', 'Duchene', 'Olivier', '1985-06-20', 'TEST1659436392IJ'),
	(5, 'test1659436401@a-pro.fr', 'Duchene', 'Olivier', '1985-06-20', 'TEST1659436401IJ'),
	(7, 'test1659436537@a-pro.fr', 'Duchene', 'Olivier', '1985-06-20', 'TEST1659436537IJ'),
	(8, 'contact@contact.com', 'Name', 'Firstname', '1985-06-20', 'MORGA657054SM9IJ'),
	(9, 'contact1@contact.com', 'Name', 'Firstname', '1985-06-20', 'MORGA657051SM9IJ');
/*!40000 ALTER TABLE "customers" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
