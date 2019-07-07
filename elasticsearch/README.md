## Elasticsearch

```
groupadd -g 1000 elasticsearch
useradd -g 1000 -u 1000 -d /home/elasticsearch -m elasticsearch
mkdir -p /home/elasticsearch/plugins/ik
cd /home/elasticsearch/plugins/ik
wget https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.2.0/elasticsearch-analysis-ik-7.2.0.zip
unzip elasticsearch-analysis-ik-7.2.0.zip
rm elasticsearch-analysis-ik-7.2.0.zip
cd /home
chown -R elasticsearch:elasticsearch elasticsearch/
cd elasticsearch/
docker-compose up -d
```
