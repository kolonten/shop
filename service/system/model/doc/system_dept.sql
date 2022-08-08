/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 127.0.0.1:3306
 Source Schema         : gogo

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 05/07/2022 15:56:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_dept
-- ----------------------------
DROP TABLE IF EXISTS `system_dept`;
CREATE TABLE `system_dept`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '机构名称',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '上级机构ID，一级机构为0',
  `sort` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '机构管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_dept
-- ----------------------------
INSERT INTO `system_dept` VALUES (1, '轻尘集团', 0, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (2, '牧尘集团', 0, 1, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (3, '三国集团', 0, 2, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (4, '上海分公司', 2, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (5, '北京分公司', 1, 1, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (6, '北京分公司', 2, 1, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (7, '技术部', 5, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (8, '技术部', 4, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (9, '技术部', 6, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (10, '市场部', 5, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (11, '市场部', 6, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (12, '魏国', 3, 0, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (13, '蜀国', 3, 1, '', '2022-07-05 15:50:38', NULL, NULL);
INSERT INTO `system_dept` VALUES (14, '吴国', 3, 2, '', '2022-07-05 15:50:38', NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
