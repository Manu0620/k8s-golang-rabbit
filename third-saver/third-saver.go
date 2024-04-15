package main

import (
    "database/sql"
    "log"

    "github.com/streadway/amqp"
    _ "github.com/lib/pq"
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
    failOnError(err, "Intento fallido de conexión a RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Intento fallido de abrir un canal")
    defer ch.Close()

    // Declarar la misma cola para el saver
    qs, err := ch.QueueDeclare(
        "queue_saver",
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
    log.Printf(" [*] Esperando mensajes. Para salir presiona CTRL+C")
    
    // Conexión a la base de datos PostgreSQL
    db, err := sql.Open("postgres", "postgres://mmadera:ivan123@postgres:5432/mydatabase?sslmode=disable")
    failOnError(err, "Intento fallido de conectar a PostgreSQL")
    defer db.Close()

    // Verificar si la tabla 'messages' existe, si no existe, crearla
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (id SERIAL PRIMARY KEY,message TEXT NOT NULL,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`)
    failOnError(err, "Intento fallido de crear la tabla 'messages'")

    for d := range msgs {
        log.Printf("Mensaje recibido por saver: %s", d.Body)

        // Insertar el mensaje en la base de datos
        _, err := db.Exec("INSERT INTO messages (message) VALUES ($1)", string(d.Body))
        failOnError(err, "Intento fallido de insertar el mensaje en la base de datos")

        log.Printf(" [✓] Mensaje guardado en la base de datos: %s", d.Body)

        err = ch.Publish(
            "",
            "queue_sender",
            false,
            false,
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(d.Body),
            },  
        )
        failOnError(err, "Fallo al publicar el mensaje")
        log.Printf(" [✓] Mensaje enviado a sender: %s", d.Body)
    }
}