package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"
	"strings"
)


var hasPing bool
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func FakeData(w http.ResponseWriter, r *http.Request)  {
	enableCors(&w)
	//
	//N, err := strconv.Atoi(r.FormValue("N"))
	//if err != nil {
	//	N = 10
	//}
	//
	//// Create N random integers
	//rand.Seed(time.Now().UTC().UnixNano())
	//var data []int
	//for i := 0; i < N; i++ {
	//	data = append(data, rand.Intn(100))
	//}
	//time.Sleep(1 * time.Second) // just for fun, pause a bit
	//// return the results
	//out, _ := json.MarshalIndent(&data, "", "  ")
	//log.Print(out)
	//fmt.Fprintf(w, string(out))
	var payload []int
	payload = append(payload, 12, 34)
	code := 200
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ping (ip string) (bool, error) {
	var pingError error
	var hasPing bool

	out, err := exec.Command("ping", ip, "-w", "10", "-i", "3").Output()
	if err != nil {
		pingError = err
	}
	pingError = nil
	fmt.Println(string(out))
	if strings.Contains(strings.ToLower(string(out)), strings.ToLower("Lost = 0")) {
		hasPing = true
	} else {
		hasPing = false
	}
	return hasPing, pingError
}


func PingData(w http.ResponseWriter, r *http.Request)  {
	enableCors(&w)
	ip := r.FormValue("ip")
	hasPing, pingError := ping(ip)
	if pingError != nil {
		code := 500
		w.WriteHeader(code)
		return
	}
	code := 200
	response, _ := json.Marshal(hasPing)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
		w.Write(response)
}

func responseTime (ip string) (string, error) {
	fmt.Println(ip)
	resp, err := http.Get(ip)
	defer resp.Body.Close()
	fmt.Println(resp)
	return "15", err
}

func ResponseTimeData(w http.ResponseWriter, r *http.Request)  {
	enableCors(&w)
	ip := r.FormValue("ip")
	time, err := responseTime(ip)
	if err != nil {
		code := 500
		w.WriteHeader(code)
		return
	}
	code := 200
	response, _ := json.Marshal(time)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data", FakeData).Methods("GET", "OPTIONS")
	r.HandleFunc("/ping", PingData).Methods("GET")
	r.HandleFunc("/responseTime", ResponseTimeData).Methods("GET")
	port := "3000"
	err := http.ListenAndServe(":"+port, r)
	fmt.Println(err)
}
