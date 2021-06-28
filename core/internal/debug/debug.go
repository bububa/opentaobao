package debug

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/bububa/opentaobao/model"
)

func PrintError(err error, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [ERROR]", err)
}

func PrintStringResponse(str string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [RESPONSE]", str)
}

func PrintGetRequest(url string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [API] GET", url)
}

func PrintPostJSONRequest(url string, data string, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	body := []byte(data)
	buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
	if err := json.Indent(buf, body, "", "    "); err == nil {
		body = buf.Bytes()
	}
	log.Printf(format, url, body)
}

func PrintPostMultipartRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [API] multipart/form-data POST", url)
}

func DecodeJSONHttpResponse(r io.Reader, v interface{}, debug bool) error {
	if !debug {
		return json.NewDecoder(r).Decode(v)
	}
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	body2 := body
	buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	if err := json.Indent(buf, body2, "", "    "); err == nil {
		body2 = buf.Bytes()
	}
	log.Printf("[DEBUG] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}

func DecodeXMLHttpResponse(r io.Reader, v interface{}, debug bool) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	var isErrResponse bool
	if !debug {
		decoder := xml.NewDecoder(bytes.NewReader(body))
		for {
			t, err := decoder.Token()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
					return err
				}
				break
			}
			t = xml.CopyToken(t)
			switch t := t.(type) {
			case xml.StartElement:
				if t.Name.Local == "error_response" {
					isErrResponse = true
				}
				break
			}
		}
	} else {
		buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
		decoder := xml.NewDecoder(bytes.NewReader(body))
		encoder := xml.NewEncoder(buf)
		encoder.Indent("", "  ")
		for {
			t, err := decoder.Token()
			if err == io.EOF {
				encoder.Flush()
				break
			}
			if err != nil {
				return err
			}
			t = xml.CopyToken(t)
			switch t := t.(type) {
			case xml.StartElement:
				if t.Name.Local == "error_response" {
					isErrResponse = true
				}
			}
			err = encoder.EncodeToken(t)
			if err != nil {
				return err
			}
		}
		log.Printf("[DEBUG] [API] http response body:\n%s\n", buf.String())
	}
	if isErrResponse {
		var errResp model.ErrorResponse
		if err := xml.Unmarshal(body, &errResp); err != nil {
			return err
		}
		return errResp
	}
	return xml.Unmarshal(body, v)
}
