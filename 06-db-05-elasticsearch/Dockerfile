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