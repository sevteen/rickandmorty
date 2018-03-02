package com.example.mortytheconsumer;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;

import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.Date;
import java.util.concurrent.CountDownLatch;

public class Listener {
    public static Logger logger = LoggerFactory.getLogger(MortyTheConsumerApplication.class);

    private final CountDownLatch latch = new CountDownLatch(3);

    @KafkaListener(topics = "test")
    public void listen(ConsumerRecord<?, ?> cr) throws Exception {
//        logger.info(cr.toString());
        System.out.println(cr.timestamp()  +" "+cr.value());
        latch.countDown();
    }
}
