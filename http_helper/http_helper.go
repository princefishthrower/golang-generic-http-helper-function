package http_helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// in the case of GET, the parameter queryParameters is transferred to the URL as query parameters
// in the case of POST, the parameter body, an io.Reader, is used
func MakeHTTPRequest[T any](fullUrl string, httpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, responseType T) (T, error) {
	client := http.Client{}
	u, err := url.Parse(fullUrl)
	if err != nil {
		return responseType, err
	}

	// if it's a GET, we need to append the query parameters.
	if httpMethod == "GET" {
		q := u.Query()

		for k, v := range queryParameters {
			// this depends on the type of api, you may need to do it for each of v
			q.Set(k, strings.Join(v, ","))
		}
		// set the query to the encoded parameters
		u.RawQuery = q.Encode()
	}

	// regardless of GET or POST, we can safely add the body
	req, err := http.NewRequest(httpMethod, u.String(), body)
	if err != nil {
		return responseType, err
	}

	// for each header passed, add the header value to the request
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// optional: log the request for easier stack tracing
	log.Printf("%s %s\n", httpMethod, req.URL.String())

	// finally, do the request
	res, err := client.Do(req)
	if err != nil {
		return responseType, err
	}

	if res == nil {
		return responseType, fmt.Errorf("error: calling %s returned empty response", u.String())
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return responseType, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return responseType, fmt.Errorf("error calling %s:\nstatus: %s\nresponseData: %s", u.String(), res.Status, responseData)
	}

	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		log.Printf("error unmarshaling response: %+v", err)
		return responseType, err
	}

	return responseObject, nil
}
