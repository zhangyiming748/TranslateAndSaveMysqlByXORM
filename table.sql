/*
 Navicat Premium Data Transfer

 Source Server         : 本地容器
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32-0ubuntu0.22.04.2)
 Source Host           : localhost:3306
 Source Schema         : baiduTranslate

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32-0ubuntu0.22.04.2)
 File Encoding         : 65001

 Date: 21/04/2023 16:56:13
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for word
-- ----------------------------
DROP TABLE IF EXISTS `word`;
CREATE TABLE `word`
(
    `zh_cn`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '中文词',
    `en_us`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '英文词',
    `id`          int NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `create_time` datetime DEFAULT NULL COMMENT '创建时间',
    `update_time` datetime DEFAULT NULL COMMENT '更新时间',
    `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_zh_0900_as_cs;

SET
FOREIGN_KEY_CHECKS = 1;
