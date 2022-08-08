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

 Date: 05/07/2022 15:57:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `site_id` int(11) NOT NULL DEFAULT 0 COMMENT '站点ID',
  `job_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '岗位ID',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `nick_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态  0：禁用   1：正常',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_user
-- ----------------------------
INSERT INTO `system_user` VALUES (1, 0, 1, 4, 'admin', '超管', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'admin@qq.com', '13612345678', 1, 'admin', 'admin', '2018-08-14 11:11:11', '2018-08-14 11:11:11', NULL);
INSERT INTO `system_user` VALUES (22, 0, 1, 7, 'liubei', '刘备', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:43:00', '2019-01-10 11:41:13', NULL);
INSERT INTO `system_user` VALUES (23, 0, 1, 7, 'zhaoyun', '赵云', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:43:44', '2018-09-23 19:43:52', NULL);
INSERT INTO `system_user` VALUES (24, 0, 1, 11, 'zhugeliang', '诸葛亮', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 7, 'admin', 'admin', '2018-09-23 19:44:23', '2018-09-23 19:44:29', NULL);
INSERT INTO `system_user` VALUES (25, 0, 1, 8, 'caocao', '曹操', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:45:32', '2019-01-10 17:59:14', NULL);
INSERT INTO `system_user` VALUES (26, 0, 1, 10, 'dianwei', '典韦', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:45:48', '2018-09-23 19:45:57', NULL);
INSERT INTO `system_user` VALUES (27, 0, 1, 8, 'xiahoudun', '夏侯惇', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:46:09', '2018-09-23 19:46:17', NULL);
INSERT INTO `system_user` VALUES (28, 0, 1, 10, 'xunyu', '荀彧', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:46:38', '2018-11-04 15:33:17', NULL);
INSERT INTO `system_user` VALUES (29, 0, 1, 10, 'sunquan', '孙权', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:46:54', '2018-09-23 19:47:03', NULL);
INSERT INTO `system_user` VALUES (30, 0, 1, 11, 'zhouyu', '周瑜', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:47:28', '2018-09-23 19:48:04', NULL);
INSERT INTO `system_user` VALUES (31, 0, 1, 11, 'luxun', '陆逊', '', '123456', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 'admin', 'admin', '2018-09-23 19:47:44', '2018-09-23 19:47:58', NULL);
INSERT INTO `system_user` VALUES (32, 0, 1, 11, 'huanggai', '黄盖', '', '', '', 'test@qq.com', '13889700023', 1, '', 'admin', '2018-09-23 19:48:38', '2021-04-03 15:43:52', NULL);
INSERT INTO `system_user` VALUES (33, 0, 1, 2, '1', '1', '', '123456', '123456', '1', '1', 1, 'admin', 'admin', '2021-04-26 17:57:50', '2021-04-26 17:57:50', NULL);
INSERT INTO `system_user` VALUES (35, 0, 1, 2, '12', '1', '', '123456', '123456', '1', '1', 1, 'admin', 'admin', '2021-04-26 18:01:53', '2021-04-26 18:01:54', NULL);
INSERT INTO `system_user` VALUES (36, 0, 1, 2, '12313', '12', '', '123456', '123456', '1', '1', 1, 'admin', 'admin', '2021-04-26 18:03:07', '2021-04-26 18:03:07', NULL);
INSERT INTO `system_user` VALUES (37, 0, 1, 3, '324', '1', '', '123456', '123456', '1', '1', 1, 'admin', 'admin', '2021-04-26 18:07:31', '2021-04-26 18:07:32', NULL);
INSERT INTO `system_user` VALUES (38, 0, 4, 7, 'aa', 'aa', '', '123456', '123456', 'a', 'a', 1, 'admin', 'admin', '2021-04-27 11:24:14', '2021-04-27 11:24:14', NULL);
INSERT INTO `system_user` VALUES (39, 0, 1, 2, '133', '133', '', '', '', '1121', '1', 1, 'admin', 'admin', '2021-04-27 12:30:15', '2021-04-27 13:53:40', NULL);
INSERT INTO `system_user` VALUES (40, 0, 4, 8, 'liu', 'liu', '', '123456', '123456', '1002219331@qq.com', '18613030352', 1, 'admin', 'admin', '2021-04-27 13:47:42', '2021-04-27 13:47:42', NULL);

SET FOREIGN_KEY_CHECKS = 1;
