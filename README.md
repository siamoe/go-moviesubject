# go-moviesubject

基于 go-douban、gorm 实现，让你一键拥有 douban 片单、豆列影视主题功能。

### 建表

拿 mysql 举个栗子：

```sql
CREATE TABLE `movie_subject`
(
    `id`            bigint                                  NOT NULL AUTO_INCREMENT COMMENT '主键',
    `subject_id`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '豆列或者片单的 id',
    `name`          varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
    `category`      varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '分类',
    `is_default`    tinyint(1) NOT NULL COMMENT '默认的',
    `display_order` int                                     NOT NULL COMMENT '排序',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='影视主题';
```

### 影视主题功能

- 初始化列表
- 获取列表
- 排序
- 添加
- 删除
- 重命名
- 恢复默认排序
- 查主题明细