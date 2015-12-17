package main
import (
	"net/http"
	"io/ioutil"
	"io"
	"fmt"
	"time"
)

func Check(url string, method string, expectedStatusCode int) Status{
	client := &http.Client{Timeout: 500 * time.Millisecond}

	response, err := client.Get(url)
	if err != nil {
		return Status {
			url,
			ERROR,
			err.Error(),
		}
	}

	defer response.Body.Close()

	if response.StatusCode != expectedStatusCode {
		return Status {
			url,
			ERROR,
			fmt.Sprintf("Expected status: %d, but Returned status: %d", expectedStatusCode, response.StatusCode),
		}
	}

	io.Copy(ioutil.Discard, response.Body)

	return Status {
		Name: url,
		Health: OK,
	}
}

func CheckAll() []Status {
	statuses := []Status{}
	statuses = append(statuses, Check("https://api.efset.org", "GET", 200))
	statuses = append(statuses, Check("https://cdn-origin.efset.org/efset-widget/v1/widget.js", "GET", 200))
	statuses = append(statuses, Check("https://cdn-origin.efset.org/athena/v1/index.js", "GET", 200))
	statuses = append(statuses, Check("https://www.efset.org", "GET", 200))
	return statuses
}
