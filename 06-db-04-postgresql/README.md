## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.

Подключитесь к БД PostgreSQL используя `psql`.

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

**Найдите и приведите** управляющие команды для:
- вывода списка БД
- подключения к БД
- вывода списка таблиц
- вывода описания содержимого таблиц
- выхода из psql

```
db=# \l
                             List of databases
   Name    | Owner | Encoding |  Collate   |   Ctype    | Access privileges
-----------+-------+----------+------------+------------+-------------------
 db        | user  | UTF8     | en_US.utf8 | en_US.utf8 |
 postgres  | user  | UTF8     | en_US.utf8 | en_US.utf8 |
 template0 | user  | UTF8     | en_US.utf8 | en_US.utf8 | =c/user          +
           |       |          |            |            | user=CTc/user
 template1 | user  | UTF8     | en_US.utf8 | en_US.utf8 | =c/user          +
           |       |          |            |            | user=CTc/user
(4 rows)

db=# \conninfo
You are connected to database "db" as user "user" via socket in "/var/run/postgresql" at port "5432".
db=# \dt
Did not find any relations.
db=# \d+
Did not find any relations.
db=# \q
root@6f43fee2e760:/#

```

## Задача 2

Используя `psql` создайте БД `test_database`.

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/virt-11/06-db-04-postgresql/test_data).

Восстановите бэкап БД в `test_database`.

Перейдите в управляющую консоль `psql` внутри контейнера.

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.

Используя таблицу [pg_stats](https://postgrespro.ru/docs/postgresql/12/view-pg-stats), найдите столбец таблицы `orders`
с наибольшим средним значением размера элементов в байтах.

**Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.

```
test_database=# analyze verbose public.orders;
INFO:  analyzing "public.orders"
INFO:  "orders": scanned 1 of 1 pages, containing 8 live rows and 0 dead rows; 8 rows in sample, 8 estimated total rows
ANALYZE
test_database=# select avg_width from pg_stats where tablename='orders';
 avg_width
-----------
         4
        16
         4
(3 rows)

```

## Задача 3

Архитектор и администратор БД выяснили, что ваша таблица orders разрослась до невиданных размеров и
поиск по ней занимает долгое время. Вам, как успешному выпускнику курсов DevOps в нетологии предложили
провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

Предложите SQL-транзакцию для проведения данной операции.

```
test_database=# begin;
BEGIN
test_database=*# create table orders_1 (check (price > 499)) inherits (orders);
CREATE TABLE
test_database=*# create table orders_2 (check (price <= 499)) inherits (orders);
CREATE TABLE
test_database=*# insert into orders_1 (select * from orders where price > 499);
INSERT 0 3
test_database=*# insert into orders_2 (select * from orders where price <= 499);
INSERT 0 5
test_database=*# commit;
COMMIT
test_database=# \dt
         List of relations
 Schema |   Name   | Type  | Owner
--------+----------+-------+-------
 public | orders   | table | user
 public | orders_1 | table | user
 public | orders_2 | table | user
(3 rows)

test_database=#

```

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?

```
Можно было бы сделать правила

create rule orders_1_more_rule as on insert to orders where ( price > 499 ) do instead insert into orders_1 values (new.*);
create rule orders_2_less_rule as on insert to orders where ( price <= 499 ) do instead insert into orders_2 values (new.*);

```

## Задача 4

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

```
Можно добавить ограничение на уровне столбца 
title character varying(80) NOT NULL UNIQUE,

Или на уровне таблицы

CREATE TABLE public.orders (
    id integer NOT NULL,
    title character varying(80) NOT NULL,
    price integer DEFAULT 0,
    UNIQUE (title)
);
```