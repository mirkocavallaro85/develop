import pika

# Connetti a RabbitMQ
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

# Definisci la coda di input
channel.queue_declare(queue='input_queue')
channel.queue_declare(queue='output_queue')

# Funzione per eseguire la somma
def add_numbers(ch, method, properties, body):
    numbers = body.decode().split(',')
    result = sum(map(int, numbers))
    print(f"Ricevuto: {numbers}, Risultato: {result}")

    # Invia il risultato a RabbitMQ
    channel.basic_publish(exchange='', routing_key='output_queue', body=str(result))

# Registra la funzione come consumatore
channel.basic_consume(queue='input_queue', on_message_callback=add_numbers, auto_ack=True)

print('In attesa di messaggi. Premi CTRL+C per uscire.')
channel.start_consuming()
