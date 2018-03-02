[start kafka on docker]

docker run -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=localhost --env ADVERTISED_PORT=9092 -d --name kafka spotify/kafka

[create topic]
kafka-topics.bat --create --topic test --zookeeper %ZOOKEEPER% --replication-factor 1 --partitions 1

[create producer for testing]
 kafka-console-producer.bat --broker-list %KAFKA% --topic test