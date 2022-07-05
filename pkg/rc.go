package pkg

import (
	"io/ioutil"
	"net/http"
)

func GetRemoteConfig(token string) ([]byte, error) {
	resp, err := http.Get(token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
