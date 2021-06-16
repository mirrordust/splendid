--' example.sql
DROP TABLE IF EXISTS "birthday";
CREATE TABLE "birthday" (
    "name" varchar(50) DEFAULT NULL,
    "born" DATETIME DEFAULT CURRENT_TIMESTAMP
);