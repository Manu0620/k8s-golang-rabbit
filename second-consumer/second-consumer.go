package main

import (
	"log"

	"github.com/streadway/amqp"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT") 
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {
	conn, err := amqp.Dial("amqp://" + rabbit_user + ":" +rabbit_password + "@" + rabbit_host + ":" + rabbit_port +"/")
	failOnError(err, "No se pudo conectar a RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "No se pudo abrir un canal")
	defer ch.Close()

	qc, err := ch.QueueDeclare(
		"queue_consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "No se pudo declarar una cola")

	msgs, err := ch.Consume(
		qc.Name,
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
		log.Printf("Mensaje recibido: %s", d.Body)
		// Aquí puedes agregar la lógica para procesar el mensaje según tus necesidades

		
		err = ch.Publish(
			"",
			"queue_saver",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(d.Body),
			},
		)
		failOnError(err, "Fallo al publicar el mensaje")
        log.Printf(" [✓] Mensaje enviado a saver: %s", d.Body)

		// Borra el mensaje de la cola
		err := ch.Cancel("", false)
		log.Printf("Mensaje eliminado: %s", d.Body)
		failOnError(err, "No se pudo cancelar el consumidor")
	
	}
}