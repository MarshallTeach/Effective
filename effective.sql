-- MySQL dump 10.13  Distrib 8.0.12, for Win64 (x86_64)
--
-- Host: localhost    Database: effective
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
-- Table structure for table `effective_basic_info`
--

DROP TABLE IF EXISTS `effective_basic_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `effective_basic_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `website_name` varchar(255) DEFAULT '' COMMENT '网站名称',
  `website_url` varchar(255) DEFAULT '' COMMENT '网站地址',
  `website_logo_url` varchar(255) DEFAULT '' COMMENT '网站logo',
  `website_favicon_url` varchar(255) DEFAULT '' COMMENT '网站favicon',
  `website_company_name` varchar(255) DEFAULT '' COMMENT '公司名称',
  `website_contact` varchar(255) DEFAULT '' COMMENT '联系人',
  `website_mobile` varchar(255) DEFAULT '' COMMENT '联系手机',
  `website_hot_line` varchar(255) DEFAULT '' COMMENT '服务热线',
  `website_tel` varchar(255) DEFAULT '' COMMENT '联系电话',
  `website_fax` varchar(255) DEFAULT '' COMMENT '传真',
  `website_qq` varchar(255) DEFAULT '' COMMENT 'qq',
  `website_email` varchar(255) DEFAULT '' COMMENT 'email',
  `website_address` varchar(255) DEFAULT '' COMMENT '公司地址',
  `website_qrcode_url` varchar(255) DEFAULT '' COMMENT '二维码',
  `website_header_msg` varchar(255) DEFAULT '' COMMENT '头部信息',
  `website_footer_msg` varchar(255) DEFAULT '' COMMENT '底部信息',
  `website_copyright` varchar(255) DEFAULT '' COMMENT '版权信息',
  `website_icp` varchar(255) DEFAULT '' COMMENT 'icp备案号',
  `website_support` varchar(255) DEFAULT '' COMMENT '技术支持',
  `website_disclaim` varchar(255) DEFAULT '' COMMENT '免责声明',
  `created_on` int(11) DEFAULT NULL,
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `deleted_on` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='网站基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `effective_basic_info`
--

LOCK TABLES `effective_basic_info` WRITE;
/*!40000 ALTER TABLE `effective_basic_info` DISABLE KEYS */;
INSERT INTO `effective_basic_info` VALUES (1,'','','','','','','','','','','','','','','','','','','','',NULL,0,NULL);
/*!40000 ALTER TABLE `effective_basic_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `effective_user`
--

DROP TABLE IF EXISTS `effective_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `effective_user` (
  `id` varchar(50) NOT NULL,
  `username` varchar(50) DEFAULT NULL COMMENT '账号',
  `password` varchar(50) DEFAULT NULL COMMENT '密码',
  `role` int(11) DEFAULT NULL COMMENT '角色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `effective_user`
--

LOCK TABLES `effective_user` WRITE;
/*!40000 ALTER TABLE `effective_user` DISABLE KEYS */;
INSERT INTO `effective_user` VALUES ('a65e6333-3b6b-11ea-b019-80e82c926305','zhuwei','fb8c0a7d404832d266baa760e60b83de',NULL);
/*!40000 ALTER TABLE `effective_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-01-23 10:28:44
