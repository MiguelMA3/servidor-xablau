package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const storagePath = "./storage/"

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Servidor S3 simulado iniciado na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	filename := getFilenameFromRequest(r)
	filepath := filepath.Join(storagePath, filename)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		http.Error(w, "Arquivo não encontrado", http.StatusNotFound)
		return
	}
	w.Write(data)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	filename := getFilenameFromRequest(r)
	filepath := filepath.Join(storagePath, filename)
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Erro ao ler o arquivo", http.StatusInternalServerError)
		return
	}
	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		http.Error(w, "Erro ao salvar o arquivo", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Arquivo %s salvo com sucesso", filename)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	filename := getFilenameFromRequest(r)
	filepath := filepath.Join(storagePath, filename)
	err := os.Remove(filepath)
	if err != nil {
		http.Error(w, "Erro ao excluir o arquivo", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Arquivo %s excluído com sucesso", filename)
}

func getFilenameFromRequest(r *http.Request) string {
	return r.URL.Path[1:]
}
