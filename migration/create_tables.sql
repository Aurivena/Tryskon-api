CREATE TABLE "Organization"
(
    uuid uuid PRIMARY KEY unique,
	name varchar,
	tag varchar
);

CREATE TABLE "User"  (
	uuid uuid PRIMARY KEY unique,
	login varchar unique,
	email varchar null unique,
	password_hash varchar,
	description varchar null,
	status int
);

CREATE TABLE "Task"  (
	uuid uuid PRIMARY KEY unique,
	title varchar,
	description varchar null,
	tag varchar unique,
	complexity int,
	is_done bool default(false)
);

CREATE TABLE "Status" (
	id serial PRIMARY KEY unique,
	status varchar default('Inactive')
);

CREATE TABLE "Complexity"  (
	id serial PRIMARY KEY  unique,
	complexity varchar default('Ease')
);

CREATE TABLE "TasksUser"  (
	"user" uuid,
	task uuid,
    PRIMARY KEY("user",task)
);

CREATE TABLE "UserOrganization"  (
	"user" uuid unique,
	organization uuid unique,
    PRIMARY KEY("user",organization)
);

ALTER TABLE "UserOrganization"
    ADD CONSTRAINT "UserOrganization_fk1" Foreign Key ("user") REFERENCES "User" ("uuid");

ALTER TABLE "UserOrganization"
    ADD CONSTRAINT "UserOrganization_fk0" Foreign Key ("organization") REFERENCES "Organization" ("uuid");

ALTER TABLE "TasksUser"
    ADD CONSTRAINT "TasksUser_fk0" Foreign Key ("user") REFERENCES "User" ("uuid");

ALTER TABLE "TasksUser"
    ADD CONSTRAINT "TasksUser_fk1" Foreign Key ("task") REFERENCES "Task" ("uuid");

ALTER TABLE "Task"
    ADD CONSTRAINT "Task_fk0" Foreign Key ("complexity") REFERENCES "Complexity" ("id");

ALTER TABLE "User"
    ADD CONSTRAINT "User_fk0" Foreign Key ("status") REFERENCES "Status" ("id");


/* Индексы для таблицы User */
CREATE INDEX idx_user_login_email ON "User" (login,email);

/* Индексы для таблицы Task */
CREATE INDEX idx_task_user ON "Task" (title);


