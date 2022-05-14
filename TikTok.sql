CREATE DATABASE IF NOT EXISTS `TikTok` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `TikTok`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`        varchar(128)        NOT NULL DEFAULT '' COMMENT '用户昵称',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `follow_id` bigint(20) unsigned NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='关注表';

DROP TABLE IF EXISTS `follower`;
CREATE TABLE `follower`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `follower_id` bigint(20) unsigned NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='粉丝表';

DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `video_id` bigint(20) unsigned NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='点赞表';

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `video_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `use_id` bigint(20) unsigned NOT NULL,
    `comtent` varchar(128) NOT NULL,
    PRIMARY KEY (`video_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`
(
    `video_id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`     bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `play_url`       varchar(128)        NOT NULL,
    `cover_url`     varchar(128)         NOT NULL,
    `create_time` timestamp           NOT NULL,
    PRIMARY KEY (`video_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='视频表';

DROP TABLE IF EXISTS `password`;
CREATE TABLE `password`
(
    `user_id`     bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `user_name`       varchar(128)        NOT NULL,
    `password`     varchar(128)         NOT NULL,
    `token` varchar(128)           NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='登录密码表';