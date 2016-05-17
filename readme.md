# Go HTTP Proxy

## Usage

Server side:
`./server -port 8888`

Client side:
`./client -remote serveraddr:8888`

## Note

Browser ---HTTP Request---> Client(located at local) ---HTTPS Security Request---> Server(em... might at foriegn)