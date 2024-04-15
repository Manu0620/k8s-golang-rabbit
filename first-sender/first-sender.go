package main

import (
    "log"
    "github.com/streadway/amqp"
    "github.com/google/uuid"
    "os" //e Importar el paquete uuid
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func generateRandomMessage() string {
    // Generar un UUID aleatorio como mensaje
    uuid := uuid.New()
    message := "Mensaje: " + uuid.String()
    return message
}

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT") 
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {
    conn, err := amqp.Dial("amqp://" + rabbit_user + ":" +rabbit_password + "@" + rabbit_host + ":" + rabbit_port +"/")
    failOnError(err, "Intento fallido de conexión a RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Intento fallido de abrir un canal")
    defer ch.Close()

    body := generateRandomMessage()
    
    err = ch.Publish(
        "",
        "queue_consumer",
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
    failOnError(err, "Fallo al publicar el mensaje")
    log.Printf(" [✓] Mensaje enviado a consumer: %s", body)

    qs, err := ch.QueueDeclare(
        "queue_sender",
        false,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Intento fallido de declarar una cola")

    msgs, err := ch.Consume(
        qs.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "No se pudo registrar un consumidor")
	log.Printf(" [*] Esperando mensajes. Para salir presiona CTRL+C")

    for d := range msgs {
        log.Printf("Mensaje recibido por sender desde saver: %s", d.Body)
    }
}