-- --------------------------------------------------------
-- 主机:                           192.168.9.9
-- 服务器版本:                        8.0.19 - MySQL Community Server - GPL
-- 服务器操作系统:                      macos10.15
-- HeidiSQL 版本:                  11.1.0.6147
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 gormgendemo 的数据库结构
CREATE DATABASE IF NOT EXISTS `gormgendemo` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `gormgendemo`;

-- 导出  表 gormgendemo.address 结构
CREATE TABLE IF NOT EXISTS `address` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `uid` int unsigned NOT NULL,
    `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `city` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_time` int unsigned NOT NULL,
    `create_time` int unsigned NOT NULL,
    `delete_time` int unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `uid` (`uid`)
    ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 数据导出被取消选择。

-- 导出  表 gormgendemo.hobby 结构
CREATE TABLE IF NOT EXISTS `hobby` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `updated_at` int unsigned NOT NULL,
    `created_at` int unsigned NOT NULL,
    `deleted_at` int unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- 数据导出被取消选择。

-- 导出  表 gormgendemo.user 结构
CREATE TABLE IF NOT EXISTS `user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `age` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
    `balance` decimal(11,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '余额',
    `updated_at` datetime NOT NULL COMMENT '更新时间',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 数据导出被取消选择。

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
