-- MySQL dump 10.13  Distrib 5.5.54, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: attendance
-- ------------------------------------------------------
-- Server version	5.5.54-0ubuntu0.12.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ComputerGraphics`
--

DROP TABLE IF EXISTS `ComputerGraphics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ComputerGraphics` (
  `date` date NOT NULL DEFAULT '0000-00-00',
  `time` varchar(50) NOT NULL DEFAULT '',
  `numberOfLectures` int(11) DEFAULT NULL,
  `absenteeEid` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`date`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ComputerGraphics`
--

LOCK TABLES `ComputerGraphics` WRITE;
/*!40000 ALTER TABLE `ComputerGraphics` DISABLE KEYS */;
/*!40000 ALTER TABLE `ComputerGraphics` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DataMining`
--

DROP TABLE IF EXISTS `DataMining`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DataMining` (
  `date` date NOT NULL DEFAULT '0000-00-00',
  `time` varchar(50) NOT NULL DEFAULT '',
  `numberOfLectures` int(11) DEFAULT NULL,
  `absenteeEid` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`date`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DataMining`
--

LOCK TABLES `DataMining` WRITE;
/*!40000 ALTER TABLE `DataMining` DISABLE KEYS */;
/*!40000 ALTER TABLE `DataMining` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DistributedComputing`
--

DROP TABLE IF EXISTS `DistributedComputing`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DistributedComputing` (
  `date` date NOT NULL DEFAULT '0000-00-00',
  `time` varchar(50) NOT NULL DEFAULT '',
  `numberOfLectures` int(11) DEFAULT NULL,
  `absenteeEid` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`date`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DistributedComputing`
--

LOCK TABLES `DistributedComputing` WRITE;
/*!40000 ALTER TABLE `DistributedComputing` DISABLE KEYS */;
/*!40000 ALTER TABLE `DistributedComputing` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `InformationSecurity`
--

DROP TABLE IF EXISTS `InformationSecurity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `InformationSecurity` (
  `date` date NOT NULL DEFAULT '0000-00-00',
  `time` varchar(50) NOT NULL DEFAULT '',
  `numberOfLectures` int(11) DEFAULT NULL,
  `absenteeEid` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`date`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `InformationSecurity`
--

LOCK TABLES `InformationSecurity` WRITE;
/*!40000 ALTER TABLE `InformationSecurity` DISABLE KEYS */;
/*!40000 ALTER TABLE `InformationSecurity` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cloudComputing`
--

DROP TABLE IF EXISTS `cloudComputing`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cloudComputing` (
  `date` date NOT NULL DEFAULT '0000-00-00',
  `time` varchar(50) NOT NULL DEFAULT '',
  `numberOfLectures` int(11) DEFAULT NULL,
  `absenteeEid` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`date`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cloudComputing`
--

LOCK TABLES `cloudComputing` WRITE;
/*!40000 ALTER TABLE `cloudComputing` DISABLE KEYS */;
/*!40000 ALTER TABLE `cloudComputing` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `enrolled`
--

DROP TABLE IF EXISTS `enrolled`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `enrolled` (
  `sid` varchar(20) NOT NULL DEFAULT '',
  `eid` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`eid`,`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `enrolled`
--

LOCK TABLES `enrolled` WRITE;
/*!40000 ALTER TABLE `enrolled` DISABLE KEYS */;
/*!40000 ALTER TABLE `enrolled` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `faculty`
--

DROP TABLE IF EXISTS `faculty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `faculty` (
  `fid` varchar(20) NOT NULL,
  `fname` varchar(50) DEFAULT NULL,
  `contact_number` int(10) DEFAULT NULL,
  `email_id` varchar(25) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`fid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `faculty`
--

LOCK TABLES `faculty` WRITE;
/*!40000 ALTER TABLE `faculty` DISABLE KEYS */;
/*!40000 ALTER TABLE `faculty` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student` (
  `eid` varchar(20) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `contact_number` int(10) DEFAULT NULL,
  `email_id` varchar(25) DEFAULT NULL,
  `semester` varchar(10) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`eid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--

LOCK TABLES `student` WRITE;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
/*!40000 ALTER TABLE `student` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subject_table`
--

DROP TABLE IF EXISTS `subject_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subject_table` (
  `sid` varchar(50) NOT NULL,
  `sname` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subject_table`
--

LOCK TABLES `subject_table` WRITE;
/*!40000 ALTER TABLE `subject_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `subject_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subject_taught`
--

DROP TABLE IF EXISTS `subject_taught`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subject_taught` (
  `fid` varchar(20) NOT NULL DEFAULT '',
  `sid` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`fid`,`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subject_taught`
--

LOCK TABLES `subject_taught` WRITE;
/*!40000 ALTER TABLE `subject_taught` DISABLE KEYS */;
/*!40000 ALTER TABLE `subject_taught` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-04-02 20:09:03
