import pika

# Connetti a RabbitMQ
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

# Definisci la coda di input
channel.queue_declare(queue='input_queue')

# Messaggio da inviare (due numeri separati da virgola)
message = "100,7"

# Invia il messaggio alla coda di input
channel.basic_publish(exchange='', routing_key='input_queue', body=message)

print(f"Inviato il messaggio: {message}")

# Chiudi la connessione
connection.close()
