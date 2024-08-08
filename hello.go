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

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		menu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa.")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido!")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Bruno"
	versao := 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println()
}

func lerComando() int { //retorno é o int
	var comando int
	fmt.Print("Escolha uma opção: ")
	fmt.Scan(&comando)

	fmt.Println("\nO comando escolhido foi", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("------ Teste", i+1, "------")
		for _, site := range sites {
			testaSites(site)
		}
		fmt.Println()
		time.Sleep(delay * time.Second)
	}
}

func testaSites(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi acessado corretamente!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas! Status code: ", response.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha) //remove quebra de linhas e espaço

		sites = append(sites, linha)

		if err == io.EOF { //final do arquivo
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))
}
