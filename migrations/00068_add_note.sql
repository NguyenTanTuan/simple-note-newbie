-- +goose Up

CREATE TABLE IF NOT EXISTS `note` (
                                      `id` INT NOT NULL AUTO_INCREMENT,
                                      `text` varchar(255) NOT NULL,
    `description` TEXT,
    PRIMARY KEY (`id`)
    ) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS `sub_note` (
                                          `id` INT NOT NULL AUTO_INCREMENT,
                                          `fk_note` INT NOT NULL,
                                          `text` varchar(255) NOT NULL,
    `description` TEXT,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`fk_note`) REFERENCES `note`(`id`) ON DELETE NO ACTION
    ) ENGINE=INNODB;
