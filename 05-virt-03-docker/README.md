## Задача 1

Сценарий выполения задачи:

- создайте свой репозиторий на https://hub.docker.com;
- выберете любой образ, который содержит веб-сервер Nginx;
- создайте свой fork образа;
- реализуйте функциональность:
  запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.

https://hub.docker.com/r/lacri/virt-21

## Задача 2

Посмотрите на сценарий ниже и ответьте на вопрос:
"Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

--

Сценарий:

- Высоконагруженное монолитное java веб-приложение;
- Nodejs веб-приложение;
- Мобильное приложение c версиями для Android и iOS;
- Шина данных на базе Apache Kafka;
- Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;
- Мониторинг-стек на базе Prometheus и Grafana;
- MongoDB, как основное хранилище данных для java-приложения;
- Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.

```
Почти во всех кейсах я бы использовал докер. 
java приложение можно развернуть в контейнерах, для распределения нагрузки использовать балансироващик.
Node, kafka, prometheus, grafna, mongo также удобно использовать докер, есть офоициальные образы на хабе.
elasticsearch также докер, удобен для кластеризации.
Под мобильным приложением не понял что конкретно имеется ввиду, бэк, база, хранилище артифактов или какая-то тестовая среда, но почти уверен что контейнер отлично подойдет. =)
```

## Задача 3

- Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;
- Добавьте еще один файл в папку ```/data``` на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.

```bash
D:\repos\VIRT-21\05-virt-03-docker>docker run --rm -itd --name centos -v %cd%/data:/data centos:latest
7664d23430b369e1f4615286e6521f5409a44d596f432b2e443926a62a6f9812

D:\repos\VIRT-21\05-virt-03-docker>docker run --rm -itd --name debian -v %cd%/data:/data debian:latest
4ba20a0178280e6d968a220064df7def97dd51c0ea43a2ef4f971399a0052e14

D:\repos\VIRT-21\05-virt-03-docker>docker exec -it centos bash
[root@7664d23430b3 /]# ls /data
host.txt
[root@7664d23430b3 /]# ls /data
[root@7664d23430b3 /]# echo "Hello World!" > /data/centos.txt
[root@7664d23430b3 /]#
D:\repos\VIRT-21\05-virt-03-docker>docker exec -it debian bash
root@4ba20a017828:/# ls /data
centos.txt  host.txt
root@4ba20a017828:/#

```