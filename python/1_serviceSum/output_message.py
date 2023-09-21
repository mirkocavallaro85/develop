import pika

# Connetti a RabbitMQ
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

# Definisci la coda di output
channel.queue_declare(queue='output_queue')

# Funzione per gestire i messaggi di risposta
def handle_response(ch, method, properties, body):
    result = body.decode()
    print(f"Ricevuta risposta: {result}")

# Registra la funzione come consumatore per la coda di output
channel.basic_consume(queue='output_queue', on_message_callback=handle_response, auto_ack=True)

print('In attesa di risposta...')
channel.start_consuming()
