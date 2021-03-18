CREATE DATABASE IF NOT EXISTS blog_service DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS blog_tag(
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(100) DEFAULT '',
    state tinyint(3) unsigned DEFAULT '1' COMMENT '0: disable 1: enable',
    create_time int(10) unsigned DEFAULT '0',
    created_by varchar(100) DEFAULT '',
    update_time int(10) unsigned DEFAULT '0',
    updated_by varchar(100) DEFAULT '',
    delete_time int(10) unsigned DEFAULT '0',
    is_del tinyint(3) unsigned DEFAULT '0' COMMENT '0: not deleted 1: deleted',
    PRIMARY KEY (id)
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS blog_auth(
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    app_key varchar(20) DEFAULT '',
    app_secret varchar(50) DEFAULT '',
    create_time int(10) unsigned DEFAULT '0',
    created_by varchar(100) DEFAULT '',
    update_time int(10) unsigned DEFAULT '0',
    updated_by varchar(100) DEFAULT '',
    delete_time int(10) unsigned DEFAULT '0',
    is_del tinyint(3) unsigned DEFAULT '0' COMMENT '0: not deleted 1: deleted',
    PRIMARY KEY (id)
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS blog_article_tag(
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    article_id int(11) NOT NULL,
    tag_id int(10) unsigned NOT NULL DEFAULT '0',
    create_time int(10) unsigned DEFAULT '0',
    created_by varchar(100) DEFAULT '',
    update_time int(10) unsigned DEFAULT '0',
    updated_by varchar(100) DEFAULT '',
    delete_time int(10) unsigned DEFAULT '0',
    is_del tinyint(3) unsigned DEFAULT '0' COMMENT '0: not deleted 1: deleted',
    PRIMARY KEY (id)
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS blog_article(
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(100) DEFAULT '',
    description varchar(255) DEFAULT '',
    cover_image_url varchar(255) DEFAULT '',
    content longtext,
    state tinyint(3) unsigned DEFAULT '1' COMMENT '0: disable 1: enable',
    create_time int(10) unsigned DEFAULT '0',
    created_by varchar(100) DEFAULT '',
    update_time int(10) unsigned DEFAULT '0',
    updated_by varchar(100) DEFAULT '',
    delete_time int(10) unsigned DEFAULT '0',
    is_del tinyint(3) unsigned DEFAULT '0' COMMENT '0: not deleted 1: deleted',
    PRIMARY KEY (id)
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB;
