## Задача 1

В этом задании вы потренируетесь в:
- установке elasticsearch
- первоначальном конфигурировании elastcisearch
- запуске elasticsearch в docker

Используя докер образ [centos:7](https://hub.docker.com/_/centos) как базовый и
[документацию по установке и запуску Elastcisearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html):

- составьте Dockerfile-манифест для elasticsearch
- соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
- запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:
- данные `path` должны сохраняться в `/var/lib`
- имя ноды должно быть `netology_test`

В ответе приведите:
- текст Dockerfile манифеста
- ссылку на образ в репозитории dockerhub
- ответ `elasticsearch` на запрос пути `/` в json виде

https://hub.docker.com/layers/lacri/virt-21/06-db-05-elasticsearch/images/sha256-2122c18048e6d2f279e8df9405006586932d3589baa16bf8d9a1c68bb34049c1?context=repo

```dockerfile
FROM centos:7

EXPOSE 9200 9300

USER 0

COPY elasticsearch .

RUN export ES_HOME="/var/lib/elasticsearch" && \
    sha512sum -c elasticsearch-8.5.2-linux-x86_64.tar.gz.sha512 && \
    tar -xzf elasticsearch-8.5.2-linux-x86_64.tar.gz && \
    rm -f elasticsearch-8.5.2-linux-x86_64.tar.gz* && \
    mv elasticsearch-8.5.2 ${ES_HOME} && \
    useradd -m -u 1000 elasticsearch && \
    chown elasticsearch:elasticsearch -R ${ES_HOME}

COPY --chown=elasticsearch:elasticsearch config/* /var/lib/elasticsearch/config/

USER 1000

ENV ES_HOME="/var/lib/elasticsearch" \
    ES_PATH_CONF="/var/lib/elasticsearch/config"

WORKDIR ${ES_HOME}

CMD ["sh", "-c", "${ES_HOME}/bin/elasticsearch"]
```

```json
{
"name": "netology_test",
"cluster_name": "elasticsearch",
"cluster_uuid": "CqBaTEnBTyixz0FZP_b5tQ",
"version": {
"number": "8.5.2",
"build_flavor": "default",
"build_type": "tar",
"build_hash": "a846182fa16b4ebfcc89aa3c11a11fd5adf3de04",
"build_date": "2022-11-17T18:56:17.538630285Z",
"build_snapshot": false,
"lucene_version": "9.4.1",
"minimum_wire_compatibility_version": "7.17.0",
"minimum_index_compatibility_version": "7.0.0"
},
"tagline": "You Know, for Search"
}
```

Подсказки:
- возможно вам понадобится установка пакета perl-Digest-SHA для корректной работы пакета shasum
- при сетевых проблемах внимательно изучите кластерные и сетевые настройки в elasticsearch.yml
- при некоторых проблемах вам поможет docker директива ulimit
- elasticsearch в логах обычно описывает проблему и пути ее решения

Далее мы будем работать с данным экземпляром elasticsearch.

## Задача 2

В этом задании вы научитесь:
- создавать и удалять индексы
- изучать состояние кластера
- обосновывать причину деградации доступности данных

Ознакомтесь с [документацией](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html)
и добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

| Имя | Количество реплик | Количество шард |
|-----|-------------------|-----------------|
| ind-1| 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.

GET /_cat/indices

```text
green  open ind-1 fxZ6_J2cQA28RYEVuQ_O9w 1 0 0 0 225b 225b
yellow open ind-3 Qp3ditROQauyAuLHq3GwSA 4 2 0 0 900b 900b
yellow open ind-2 NZtIEvOITPmqFOFCNr11hQ 2 1 0 0 450b 450b
```

Получите состояние кластера `elasticsearch`, используя API.

GET /_cluster/health

```json
{
  "cluster_name": "elasticsearch",
  "status": "yellow",
  "timed_out": false,
  "number_of_nodes": 1,
  "number_of_data_nodes": 1,
  "active_primary_shards": 8,
  "active_shards": 8,
  "relocating_shards": 0,
  "initializing_shards": 0,
  "unassigned_shards": 10,
  "delayed_unassigned_shards": 0,
  "number_of_pending_tasks": 0,
  "number_of_in_flight_fetch": 0,
  "task_max_waiting_in_queue_millis": 0,
  "active_shards_percent_as_number": 44.44444444444444
}
```

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?

```text
В состоянии yellow у нас индексы с ненулевым количеством реплик, в кластере у нас только одна нода. 
Реплики некуда привязать
```

Удалите все индексы.

DELETE /ind-{1,2,3}

**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард,
иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

## Задача 3

В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

Используя API [зарегистрируйте](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html#snapshots-register-repository)
данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

PUT /_snapshot/netology_backup

```json
{
    "type": "fs",
    "settings": {
        "location": "/var/lib/elastisearch/snapshots"
    }
}
```

```json
{"acknowledged":true}
```

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

```text
green open test 0Ck7JLHMQOmKQ2yu0y65cA 1 0 0 0 225b 225b
```

[Создайте `snapshot`](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html)
состояния кластера `elasticsearch`.

**Приведите в ответе** список файлов в директории со `snapshot`ами.

```shell
sh-4.2$ ls -1 /var/lib/elasticsearch/snapshots
index-0
index.latest
indices
meta-orF8wsUsQDip9TtdoyNAMw.dat
snap-orF8wsUsQDip9TtdoyNAMw.dat
```

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

```text
green open test-2 KaVRkPNISb-Y5GPzkS965g 1 0 0 0 225b 225b
```

[Восстановите](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-restore-snapshot.html) состояние
кластера `elasticsearch` из `snapshot`, созданного ранее.

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

POST /_snapshot/netology_backup/snapshot_1/_restore

```json
{"accepted":true}
```

```text
green open test-2 KaVRkPNISb-Y5GPzkS965g 1 0 0 0 225b 225b
green open test   a9CRRNEAQFO0pZKggEFm0w 1 0 0 0 225b 225b
```