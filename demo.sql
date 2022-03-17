/*
 Navicat Premium Data Transfer

 Source Server         : Mac
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : localhost:3306
 Source Schema         : demo

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 17/03/2022 20:12:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for d
-- ----------------------------
DROP TABLE IF EXISTS `d`;
CREATE TABLE `d` (
  `id` int(11) NOT NULL,
  `content` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- ----------------------------
-- Records of d
-- ----------------------------
BEGIN;
INSERT INTO `d` VALUES (1, 'd1');
INSERT INTO `d` VALUES (2, 'd2');
INSERT INTO `d` VALUES (3, 'd3');
INSERT INTO `d` VALUES (4, 'd4');
COMMIT;

-- ----------------------------
-- Table structure for d_demo
-- ----------------------------
DROP TABLE IF EXISTS `d_demo`;
CREATE TABLE `d_demo` (
  `id` int(11) NOT NULL,
  `d_id` int(11) DEFAULT NULL,
  `demo_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- ----------------------------
-- Records of d_demo
-- ----------------------------
BEGIN;
INSERT INTO `d_demo` VALUES (1, 1, 1);
INSERT INTO `d_demo` VALUES (2, 1, 2);
INSERT INTO `d_demo` VALUES (3, 1, 3);
INSERT INTO `d_demo` VALUES (4, 2, 1);
INSERT INTO `d_demo` VALUES (5, 2, 2);
INSERT INTO `d_demo` VALUES (6, 3, 2);
INSERT INTO `d_demo` VALUES (7, 4, 3);
COMMIT;

-- ----------------------------
-- Table structure for demo
-- ----------------------------
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `jointime` datetime DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- ----------------------------
-- Records of demo
-- ----------------------------
BEGIN;
INSERT INTO `demo` VALUES (1, 'test_name', '2020-09-22 16:23:45', '2020-09-22 16:24:06');
INSERT INTO `demo` VALUES (2, 'test_name_2', '2021-08-09 16:36:26', '2021-08-09 16:36:32');
INSERT INTO `demo` VALUES (3, 'test_name_3', '2021-08-09 16:36:58', '2021-08-09 16:37:04');
COMMIT;

-- ----------------------------
-- Table structure for demo_details
-- ----------------------------
DROP TABLE IF EXISTS `demo_details`;
CREATE TABLE `demo_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `demo_id` int(11) DEFAULT NULL,
  `content` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- ----------------------------
-- Records of demo_details
-- ----------------------------
BEGIN;
INSERT INTO `demo_details` VALUES (1, 1, 't11');
INSERT INTO `demo_details` VALUES (2, 1, 't12');
INSERT INTO `demo_details` VALUES (3, 2, 't23');
INSERT INTO `demo_details` VALUES (4, 2, 't24');
INSERT INTO `demo_details` VALUES (5, 3, 't35');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
