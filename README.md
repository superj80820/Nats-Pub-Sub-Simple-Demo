# Nats-Pub-Sub-Simple-Demo

此Repository使用Nats來當作Pub/Sub的Server，並用Golang對Nats做了簡單的Pub/Sub Demo

## 需要安裝

* docker
* docker-compose

## 怎麼使用?

```bash
$ docker-compose up
```

會看到Golang對Nats的foo topic publish了一個Hello World訊息，並且從Nats端接收回來

![](https://imgur.com/VWewcMI.jpg)