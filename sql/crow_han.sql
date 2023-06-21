/*
 Navicat Premium Data Transfer

 Source Server         : 监控平台-dev
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 10.10.239.136:3306
 Source Schema         : crow_qin

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 07/04/2023 17:02:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(24) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `created_at` datetime(0) DEFAULT NULL,
  `updated_at` datetime(0) DEFAULT NULL,
  `deleted_at` datetime(0) DEFAULT NULL,
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `phone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `role` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pjm_project_id_IDX`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (0, 'admin', '2023-02-21 19:40:51', '2023-02-28 17:38:59', NULL, '413c971e8ce77e8b66da71a84ca0bdbe', NULL, NULL, 'admin');

SET FOREIGN_KEY_CHECKS = 1;
