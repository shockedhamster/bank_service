CREATE TABLE operation_type
(
    id serial primary key,
    type_name varchar(255) not null
);

CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE accounts
(
    id         serial primary key,
    balance    int       not null default 0,
    created    timestamp not null default now()
);

CREATE TABLE operations
(
    id             serial primary key,
    account_id     int          not null,
    amount         int          not null,
    operation_type int not null,
    created        timestamp    not null default now(),
    foreign key (account_id) references accounts (id),
    foreign key (operation_type) references operation_type (id)
);

INSERT INTO operation_type (type_name) VALUES ('deposit');
INSERT INTO operation_type (type_name) VALUES ('withdraw');
INSERT INTO operation_type (type_name) VALUES ('transfer_from');
INSERT INTO operation_type (type_name) VALUES ('transfer_to');