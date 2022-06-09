package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"time"

	_ "github.com/markcheno/go-quote"
	_ "github.com/markcheno/go-talib"
	_ "practice_go/lib"
)

const (
	c1 = iota
	c2 = iota
	c3 = iota
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type S struct{}
type Person struct {
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames,omitempty"`
	S         *S       `json:"S,omitempty"`
}

func logProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

// func (p Person) MarshalJSON() ([]byte, error) {
// 	v, err := json.Marshal(&struct{
// 		Name string
// 	}{
// 		Name: "Mr." + p.Name,
// 	})
// 	return v, err
// }

var DB = map[string]string{
	"UserKey1": "User1Secret",
	"UserKey2": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

func main() {
	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(lib.Avarage(s))

	// spy, _ := quote.NewQuoteFromYahoo("spy", "2022-01-01", "2022-06-01", quote.Daily, true)
	// fmt.Print(spy.CSV())
	// rsi2 := talib.Rsi(spy.Close, 2)
	// fmt.Println(rsi2)

	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	i := []int{5, 3, 2, 8, 7}
	s := []string{"e", "f", "j", "k", "x"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)

	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go logProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("####################")

	base, _ := url.Parse("https://example.com/")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	// req, _ := http.NewRequest("GET", endpoint, nil)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	b := []byte(`{"name":"mike","age":20,"nicknames":["a", "b", "c"]}`)
	var person Person
	err := json.Unmarshal(b, &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person.Name, person.Age, person.Nicknames)

	v, _ := json.Marshal(person)
	fmt.Println(string(v))

	const API_KEY = "UserKey1"
	const API_SECRET = "User1Secret"

	data := []byte("data")
	mac := hmac.New(sha256.New, []byte(API_SECRET))
	mac.Write(data)
	sign := hex.EncodeToString(mac.Sum(nil))

	Server(API_KEY, sign, data)

}
