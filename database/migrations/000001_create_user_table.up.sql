create table if not exists users (
  id int unique not null auto_increment primary key,
  name varchar(255) unique not null,
  email varchar(255) unique not null,
  cognito_uuid varchar(255),
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  deleted_at timestamp null
);