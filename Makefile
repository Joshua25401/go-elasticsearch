run-elasticsearch:
	docker network create elastic
	docker run --name sample-elastic --net elastic -p 9200:9200 -itd -m 1GB -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:8.10.4

reproduce-password: 
	docker exec -it sample-elastic /usr/share/elasticsearch/bin/elasticsearch-reset-password -u elastic

reproduce-enrollment:
	docker exec -it sample-elastic /usr/share/elasticsearch/bin/elasticsearch-create-enrollment-token -s kibana

test-elastic-server: 
	curl --cacert http_ca.crt -u elastic:$(ELASTIC_PASSWORD) https://localhost:9200

copy-cert:
	docker cp sample-elastic:/usr/share/elasticsearch/config/certs/http_ca.crt .

tidy-elasticsearch:
	docker container stop sample-elastic 
	docker container rm sample-elastic
	docker network rm elastic
	


