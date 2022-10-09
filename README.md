# INTRODUCTION

This tool is used to gather Internet metrics and expose it to Prometheus. Main metrics, for now, are:

* Availability
* Speed

The tool have two parts, a **client** and a **server**. The server listen in a specific port (HTTP) and can receive 
requests from client to upload a file. Both metrics (availability and speed) are derived from uploading the file.

The client side runs as a daemon and periodically trys to connect and upload a file to the server and expose both 
metrics to be scrapped by prometheus. 

# Running

## Simple Run

By default, the server side will start to run and listening on port 8080 and will accept files up to 5MiB.   

```bash
netmon start server 
```

On client side (from the network you want to test), you should run:

```bash
netmon start client --remote_server_address <ip> 
```

It will start the client side and upload a file to server IP. Client will try upload a file every 5min. 

## Configuration

### Server

The server accepts the following flags:

* `port`: Port that server will start to listen (optional, default: 8080)
* `max_file_mb`: Max file size that servers will accept from client, in MB (optional, default: 5MB)

### Client

* `remote_server_address`: Server address that will receive the file and connection (required)
* `remote_server_port`: Remote server port to connect and send the file (optional, default: 8080)
* `upload_freq_min`: Frequency that client daemon will upload the file to remote server (optional, default: 5)
* `upload_file_size_mb`: File size to upload to the remote server  (optional, default: 5)
