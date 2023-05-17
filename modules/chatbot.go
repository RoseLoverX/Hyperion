package modules

import (
	"encoding/json"
	"main/client"
	"net/http"
	"net/url"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

const BARD_API = "https://api.csa.codes/api/bard"

func BardHandler(m *tg.NewMessage) error {
	if m.Args() == "" {
		return EoR(m, "Please provide a query")
	}
	query := m.Args()
	bard, err := GetBard(query)
	if err != nil {
		return EoR(m, err.Error())
	}
	m.Reply(bard)
	return nil
}

func GetBard(query string) (string, error) {
	req, err := http.NewRequest("POST", BARD_API, strings.NewReader(url.Values{"query": {query}}.Encode()))
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var bard map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&bard)
	if err != nil {
		return "", err
	}
	if cont, ok := bard["content"]; ok {
		return cont.(string), nil
	}
	return "", nil
}

func init() {
	client.RegCmd("bard", BardHandler)
}
