package cotacao

import (
	"fmt"
	"github.com/fabioxgn/go-bot"
	"github.com/fabioxgn/go-bot/web"
)

const (
	URL = "http://developers.agenciaideias.com.br/cotacoes/json"
)

type Cotacao struct {
	Bovespa struct {
		Cotacao  string `json:"cotacao"`
		Variacao string `json:"variacao"`
	} `json:"bovespa"`
	Dolar struct {
		Cotacao  string `json:"cotacao"`
		Variacao string `json:"variacao"`
	} `json:"dolar"`
	Euro struct {
		Cotacao  string `json:"cotacao"`
		Variacao string `json:"variacao"`
	} `json:"euro"`
	Atualizacao string `json:"atualizacao"`
}

func cotacao(command *bot.Cmd) (msg string, err error) {
	data := &Cotacao{}
	err = web.GetJSON(URL, data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Dólar: %s (%s), Euro: %s (%s)",
		data.Dolar.Cotacao, data.Dolar.Variacao,
		data.Euro.Cotacao, data.Euro.Variacao), nil
}

func init() {
	bot.RegisterCommand(
		"cotacao",
		"Informa a cotação do Dólar e Euro.",
		"",
		cotacao)
}
