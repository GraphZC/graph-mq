# GraphMQ

A simple message queue using go language

GraphMQ is a project for cloud and communication subject

## Info

- **Author**: Tanaroeg O-Charoen 6410401078
- **Queue** use linked list to store messages from clients

### How it work?

- Producer send a message to the queue and the queue will store the message to queue linked list and return a status to the producer
- Consumer will request a message to subscribe the topic  then add client to subscribe table and wait for the message from the queue
- When the queue receive a message from the producer, the collect message service will be wake up and send the message to the consumer that subscribe to the topic

This method will make the queue can handle multiple consumer and producer, and client does not wait for the message sending to the consumer complete

## How to use

0. Install go and make

1. Build server and client

```bash
make build-server
make build-client
```

2. Run server

```bash
make run-server
```

3. Run client

```bash
make run-client
```

4. Subscribe to topic

- By GraphMQ client
  run the client type `SUB` enter and the type topic you want to subscribe
- By send request to server
  Connect to the server and send a message `SUBC topic` to subscribe to the topic

5. Send message to topic

- By GraphMQ client
  run the client type `PUB` enter and the type topic and message you want to send
- By send request to server
  Connect to the server and send a message `PUBL topic message` to send message to the topic
