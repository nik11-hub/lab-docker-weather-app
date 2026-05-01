package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
)

func main() {
	port := "8080"

	
	log.Printf("Data uruchomienia: %s", time.Now().Format(time.RFC3339))
	log.Println("Autor: Mikita Liaiko")
	log.Printf("Aplikacja nasłuchuje na porcie TCP: %s", port)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/weather", weatherHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<head><title>Weather App</title></head>
	<body>
		<h1>Wybierz lokalizację</h1>
		<form action="/weather" method="get">
			<select name="city">
				<option value="Lublin">Polska - Lublin</option>
				<option value="Wroclaw">Polska - Wrocław</option>
				<option value="Minsk">Białoruś - Mińsk</option>
				<option value="New York">USA - New York</option>
			</select>
			<br><br>
			<input type="submit" value="Sprawdź pogodę">
		</form>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "Nie wybrano miasta", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	url := fmt.Sprintf("https://wttr.in/%s?format=3", city)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Błąd pobierania pogody", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, "<h2>Pogoda dla: %s</h2>", city)
	fmt.Fprintf(w, "<p>%s</p>", string(body))
	fmt.Fprintf(w, `<br><a href="/">Powrót</a>`)
}
