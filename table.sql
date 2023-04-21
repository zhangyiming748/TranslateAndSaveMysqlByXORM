CREATE TABLE `word`
(
    `zh_cn`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '中文翻译',
    `other`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '其他语言',
    `kind`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs DEFAULT NULL COMMENT '外文语种',
    `id`          int NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `create_time` datetime                                                         DEFAULT NULL COMMENT '创建时间',
    `update_time` datetime                                                         DEFAULT NULL COMMENT '更新时间',
    `delete_time` datetime                                                         DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_zh_0900_as_cs;

SET
    FOREIGN_KEY_CHECKS = 1;
