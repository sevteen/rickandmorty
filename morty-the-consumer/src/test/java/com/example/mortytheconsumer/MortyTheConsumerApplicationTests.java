package com.example.mortytheconsumer;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.time.Instant;
import java.time.LocalDateTime;

@RunWith(SpringRunner.class)
@SpringBootTest
public class MortyTheConsumerApplicationTests {

	@Test
	public void contextLoads() {
		LocalDateTime.from(Instant.ofEpochSecond(Long.parseLong("1520002378474")));
	}

}
