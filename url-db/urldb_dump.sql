-- MySQL dump 10.13  Distrib 8.0.12, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: urldb
-- ------------------------------------------------------
-- Server version	8.0.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `long_urls`
--

DROP TABLE IF EXISTS `long_urls`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `long_urls` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(2000) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `long_urls`
--

LOCK TABLES `long_urls` WRITE;
/*!40000 ALTER TABLE `long_urls` DISABLE KEYS */;
INSERT INTO `long_urls` VALUES (1,'https://stackoverflow.com/questions/45736762/how-to-connect-dockers-with-compose-mysql-and-golang'),(2,'https://www.linkedin.com/pulse/running-java-application-mysql-linked-docker-deepak-sureshkumar/'),(3,'https://blog.csainty.com/2016/07/connecting-docker-containers.html'),(4,'https://github.com/go-sql-driver/mysql/'),(5,'https://www.randomurl.com/'),(6,'https://randomurl.com/9834759384759'),(7,'https://randomurl.com/'),(8,'https://lanre.wtf/blog/2017/04/08/testing-http-handlers-go/'),(9,'https://randomurl.com/8674665223082153551'),(10,'https://duckduckgo.com/?q=how+to+initalize+an+array+in+golang&t=ffab&atb=v41-2&ia=qa'),(11,'https://stackoverflow.com/questions/34623523/making-a-post-request-in-golang'),(12,'https://randomurl.com/98573948759'),(13,'https://randomurl.com/5577006791947779410'),(14,'https://randomurl.com/6129484611666145821'),(15,'https://randomurl.com/5577006791947779410'),(16,'https://randomurl.com/5577006791947779410'),(17,'https://randomurl.com/4278071406432343658'),(18,'https://randomurl.com/4278071406432343658'),(19,'https://longurl.com/2448397773495314005'),(20,'https://randomurl.com/2162005612667445379'),(21,'https://randomurl.com/2162005612667445379'),(22,'https://longurl.com/2882151827269023763'),(23,'https://randomurl.com/1002071167725390085'),(24,'https://randomurl.com/1002071167725390085'),(25,'https://longurl.com/3958632839207850688'),(26,'https://randomurl.com/2608450357297581469'),(27,'https://randomurl.com/2608450357297581469'),(28,'https://longurl.com/4989068523327257856'),(29,'https://randomurl.com/2763148733171029355'),(30,'https://randomurl.com/2763148733171029355'),(31,'https://longurl.com/844822227128958305'),(32,'https://randomurl.com/4961161656633552087'),(33,'https://randomurl.com/4961161656633552087'),(34,'https://longurl.com/9026624833785294014'),(35,'https://randomurl.com/2568732495613761450'),(36,'https://randomurl.com/2568732495613761450'),(37,'https://longurl.com/1856116199691502053'),(38,'https://randomurl.com/8088095155733846505'),(39,'https://randomurl.com/8088095155733846505'),(40,'https://longurl.com/7408045335831662660'),(41,'https://randomurl.com/8266413534353393396'),(42,'https://randomurl.com/8266413534353393396'),(43,'https://longurl.com/9073211106126012942'),(44,'https://randomurl.com/8882739719246786258'),(45,'https://randomurl.com/8882739719246786258'),(46,'https://longurl.com/8811083208250388774'),(47,'https://randomurl.com/7721990894574422991'),(48,'https://randomurl.com/7721990894574422991'),(49,'https://longurl.com/4648035244397424411'),(50,'https://randomurl.com/6049789512638828093'),(51,'https://randomurl.com/6049789512638828093'),(52,'https://longurl.com/433317372595269091'),(53,'https://randomurl.com/8335299538026353093'),(54,'https://randomurl.com/8335299538026353093'),(55,'https://longurl.com/1854983534200123267'),(56,'https://randomurl.com/6706774231356099976'),(57,'https://randomurl.com/6706774231356099976'),(58,'https://longurl.com/9149805907121635072'),(59,'https://randomurl.com/1393848413331972382'),(60,'https://randomurl.com/1393848413331972382'),(61,'https://longurl.com/8995299800454060681'),(62,'https://randomurl.com/813476974838412405'),(63,'https://randomurl.com/813476974838412405'),(64,'https://longurl.com/1632722541266529806'),(65,'https://randomurl.com/755005208102058701'),(66,'https://randomurl.com/755005208102058701'),(67,'https://longurl.com/1327608878994582111'),(68,'https://randomurl.com/3411447655466405061'),(69,'https://randomurl.com/3411447655466405061'),(70,'https://longurl.com/2784878599108029754');
/*!40000 ALTER TABLE `long_urls` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-09-11 21:31:41
