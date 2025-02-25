/*
 Navicat Premium Dump SQL

 Source Server         : 45.207.207.178
 Source Server Type    : MySQL
 Source Server Version : 50744 (5.7.44)
 Source Host           : 45.207.207.178:7002
 Source Schema         : cheemshappypay

 Target Server Type    : MySQL
 Target Server Version : 50744 (5.7.44)
 File Encoding         : 65001

 Date: 24/02/2025 11:40:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for merchants
-- ----------------------------
DROP TABLE IF EXISTS `merchants`;
CREATE TABLE `merchants` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` text NOT NULL COMMENT '昵称',
  `avatar` text COMMENT '头像',
  `introduction` text COMMENT '简介',
  `password` longtext NOT NULL,
  `email` longtext NOT NULL,
  `phone` text,
  `created_by` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for merchants_api
-- ----------------------------
DROP TABLE IF EXISTS `merchants_api`;
CREATE TABLE `merchants_api` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_id` bigint(20) unsigned NOT NULL,
  `apikey` text NOT NULL COMMENT 'Token',
  `callback_url` text NOT NULL COMMENT '回调url',
  `secret_key` text NOT NULL COMMENT '回调加密参数',
  `remark` text COMMENT '备注',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for merchants_meta
-- ----------------------------
DROP TABLE IF EXISTS `merchants_meta`;
CREATE TABLE `merchants_meta` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_id` bigint(20) unsigned NOT NULL,
  `ac` text COMMENT '钱包地址',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_id` bigint(20) unsigned NOT NULL,
  `no` text NOT NULL COMMENT '订单no',
  `c_no` text NOT NULL COMMENT '用户no',
  `amount` text NOT NULL COMMENT '价格',
  `coin` text NOT NULL COMMENT '数字币类型',
  `chain` text NOT NULL COMMENT '链',
  `notify_url` text COMMENT '回调url',
  `return_url` text COMMENT '前端重定向url',
  `pay_url` text COMMENT '支付地址',
  `tx_hash` text COMMENT '交易hash',
  `ac` text COMMENT '钱包地址',
  `api_key` text COMMENT '商户apikey',
  `time_out` int(11) DEFAULT NULL COMMENT '超时时间',
  `status` varchar(255) DEFAULT NULL COMMENT '订单状态',
  `notify_status` varchar(255) DEFAULT NULL COMMENT '订单回调结果',
  `remark` text COMMENT '备注',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `domain` text NOT NULL COMMENT '项目域名',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for sys_wallet
-- ----------------------------
DROP TABLE IF EXISTS `sys_wallet`;
CREATE TABLE `sys_wallet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ac` text NOT NULL COMMENT '钱包地址',
  `pri_key` text NOT NULL COMMENT '私钥',
  `mnemonic` varchar(255) NOT NULL COMMENT '助记词',
  `path` text NOT NULL COMMENT '钱包路径',
  `remark` text COMMENT '备注',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `current_index` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '当前派生索引',
  `version` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_type` varchar(191) NOT NULL,
  `user_id` varchar(191) NOT NULL,
  `nickname` text NOT NULL COMMENT '昵称',
  `avatar` text NOT NULL COMMENT '头像',
  `introduction` text COMMENT '简介',
  `password` longtext NOT NULL,
  `email` longtext NOT NULL,
  `phone` text,
  `login_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3),
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id` (`user_type`) USING BTREE,
  KEY `idx_users_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=174 DEFAULT CHARSET=utf8mb4;

INSERT INTO `user` (`id`, `user_type`, `user_id`, `nickname`, `avatar`, `introduction`, `password`, `email`, `phone`, `login_at`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '123456', '123', '1', NULL, '$2a$10$1DpQrC8M8iL7tOI3JFan7eu4Is7XgUdIXEy6zKK/wV8I3VzMqVgJe', 'ch@cheemshappy.com', NULL, '2025-02-18 06:14:40.264', NULL, '2025-02-25 16:43:53.518', NULL);
-- ----------------------------
-- Table structure for wallet
-- ----------------------------
DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `m_id` bigint(20) unsigned NOT NULL,
  `ac` text NOT NULL COMMENT '钱包地址',
  `pri_key` text NOT NULL COMMENT '私钥',
  `mnemonic` text NOT NULL COMMENT '助词器',
  `path` text NOT NULL COMMENT '钱包路径',
  `remark` text COMMENT '备注',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
