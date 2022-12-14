## Задача 1

Архитектор ПО решил проконсультироваться у вас, какой тип БД
лучше выбрать для хранения определенных данных.

Он вам предоставил следующие типы сущностей, которые нужно будет хранить в БД:

- Электронные чеки в json виде
- Склады и автомобильные дороги для логистической компании
- Генеалогические деревья
- Кэш идентификаторов клиентов с ограниченным временем жизни для движка аутенфикации
- Отношения клиент-покупка для интернет-магазина

Выберите подходящие типы СУБД для каждой сущности и объясните свой выбор.

```
- Электронные чеки в json виде

Для хранения json отлично подходят NoSQL субд, например документориентированные.
Для которых json является основными способом хранения данных.

- Генеалогические деревья
- Склады и автомобильные дороги для логистической компании

Из лекции я узнал о существовании т.н. графовых субд, не работал раньше с такими, да и звучит как разновидность связей в реляционнках.
Поэтому для хранения таких данных отлично подойдут графовые субд.

- Кэш идентификаторов клиентов с ограниченным временем жизни для движка аутенфикации

Ответ содержится в вопросе, для кэшей критически важна сокрость доступа, в купе с овж отлично подойдут nosql базы хранящие данные в озу. 
Redis, memcached например.

- Отношения клиент-покупка для интернет-магазина
Реляционные БД отлично подойдут для хранения связанных данных.
```


## Задача 2

Вы создали распределенное высоконагруженное приложение и хотите классифицировать его согласно
CAP-теореме. Какой классификации по CAP-теореме соответствует ваша система, если
(каждый пункт - это отдельная реализация вашей системы и для каждого пункта надо привести классификацию):

- Данные записываются на все узлы с задержкой до часа (асинхронная запись)

```
AP, PA-EL
```  

- При сетевых сбоях, система может разделиться на 2 раздельных кластера

```
AP, PA-EL
```  

- Система может не прислать корректный ответ или сбросить соединение

```
CP, PA-EC
```

А согласно PACELC-теореме, как бы вы классифицировали данные реализации?

## Задача 3

Могут ли в одной системе сочетаться принципы BASE и ACID? Почему?

```
Я думаю что если да, то с серьезными допущениями. ACID подразумеват согласованность данных, 
с другой стороны BASE возлагает ответственность за согласованность на разработчика, а не на субд.
```

## Задача 4

Вам дали задачу написать системное решение, основой которого бы послужили:

- фиксация некоторых значений с временем жизни
- реакция на истечение таймаута

Вы слышали о key-value хранилище, которое имеет механизм [Pub/Sub](https://habr.com/ru/post/278237/).
Что это за система? Какие минусы выбора данной системы?

```
Redis - СУБД класса NoSQL. 
Является базой резидентского типа, то есть, размещаемой в оперативной памяти. 
Работает со структурами данных «ключ - значение». 
Ориентирована на быстрое выполнение атомарных операций в нагруженных системах.
Redis вообще-то не pub-sub система, а key-value data storage. 
Но он на удивление хорошо реализует классический pub-sub и показывает замечательную производительность.

- Может использоваться как БД, так и как кэш-система или брокер сообщений.
- Данным можно присваивать Time-To-Live.
- Имеется встроенная система Pub/Sub.
- Поддерживает Master-Slave репликацию.

Из минусов
- Так как данные хранятся в ОЗУ, то имеется высокая вероятность потери данных.
- Сложности хранения объемных данных, из-за весьма ограниченного объема ОЗУ.
- Нет поддержки SQL
- Отсутствует разграничение прав доступа по пользователям
- “Из коробки” не имеет механизма консенсуса. При отказе ведущей реплики - необходимо вручную выбрать новую ведущую реплику.
```