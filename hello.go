package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Insira seu nome:")
	nomeUsuario := lerNome()
	fmt.Println("Bem-vindo", nomeUsuario+".")
	for {

		exibir()
		resposta := leituraComando()

		switch resposta {
		case 1:
			monitorandoSites()
		case 2:
			exibeLog()

		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		}
	}
}
func exibir() {
	fmt.Println("1 - Monitorar sites.")
	fmt.Println("2 - Log dos sites.")
	fmt.Println("0 - Sair.")
}
func leituraComando() int {
	var respostaLida int
	fmt.Scan(&respostaLida)
	return respostaLida
}
func monitorandoSites() {
	fmt.Println("Monitorando...")
	sites := leSite()
	for _, site := range sites {
		resp, err := http.Get(site)
		if err != nil {
			fmt.Println("Ocorreu um erro monitorando", err)
		}
		if resp.StatusCode == 200 {
			fmt.Println("O site", site, "está ativo!")
			registraLog(site, true)
		} else {
			fmt.Println("O site", site, "caiu")
			registraLog(site, false)
		}
		time.Sleep(5 * time.Second)
	}
}
func leSite() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu o erro lesite", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		} else if err == nil {
		} else {
			fmt.Println("Ocorreu o erro read string", err)
		}
	}
	arquivo.Close()
	return sites
}
func lerNome() string {
	var nomeLido string
	fmt.Scan(&nomeLido)
	return nomeLido
}
func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu o erro registra", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

// func lerLog() []string {
// 	fmt.Println("Exibindo Logs...")
// 	var logLida []string
// 	arquivo, err := os.Open("log.txt")
// 	if err != nil {
// 		fmt.Println("Ocorreu o erro exibelog", err)
// 	}
// 	leitor := bufio.NewReader(arquivo)
// 	for {
// 		linha, err := leitor.ReadString('\n')
// 		linha = strings.TrimSpace(linha)
// 		logLida = append(logLida, linha)
// 		if err == io.EOF {
// 			break
// 		} else if err == nil {
// 		} else {
// 			fmt.Println("Ocorreu o erro ler logs string", err)
// 		}
// 	}
// 	arquivo.Close()
// 	return logLida
// }
// func exibeLog() {
// 	logs := lerLog()
// 	for _, alog := range logs {
// 		fmt.Println(alog)
// 	}
// }
func exibeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro em exibe log", err)
	}
	fmt.Println(string(arquivo))
}
