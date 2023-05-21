import socket

def main():

    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    # Bind the socket to a specific address and port
    server_address = ('', 44022)  # Listen on all available network interfaces
    server_socket.bind(server_address)
    server_socket.listen(1)
    print("Listening on tcp 44022...")
    while True:
        client_socket, client_address = server_socket.accept()
        print("Client connected:", client_address)

        try:
            data = client_socket.recv(1024)
            if data:
                # Print the received data
                print(data.decode())

        except Exception as e:
            print("Error:", str(e))

        finally:
            client_socket.close()

if __name__ == '__main__':
    main()