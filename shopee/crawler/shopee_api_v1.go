package crawler

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func GetProduct(link string) Product {
	url := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage(link)
	time.Sleep(5 * time.Second)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(page.MustHTML()))
	if err != nil {
		fmt.Println("co loi")
	}
	var options [100]string
	doc.Find("button.product-variation").Each(func(i int, s *goquery.Selection) {
		options[i] = s.Text()
	})
	var category []string
	doc.Find("a.CyVtI7._2yC5g9").Each(func(i int, s *goquery.Selection) {
		category = append(category, s.Text())
	})
	ShipsFrom := ""
	doc.Find("div.OktMMO:contains(\"Gửi từ\")").Each(func(i int, s *goquery.Selection) {
		ShipsFrom = s.Find("div").Text()
	})
	Stock := ""
	doc.Find("div.OktMMO:contains(\"Kho hàng\")").Each(func(i int, s *goquery.Selection) {
		Stock = s.Find("div").Text()
	})
	Sold := doc.Find("div.HmRxgn").Text()
	//var products[] Product
	var productFinal Product
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		attr, exist := s.Attr("type")
		if attr == "application/ld+json" && exist == true {
			//fmt.Println(s.Text())
			var product Product
			json.Unmarshal([]byte(s.Text()), &product)
			if product.Description != "" {
				productFinal = product
			}
		}
	})
	productFinal.Sold, _ = strconv.Atoi(Sold)
	productFinal.Stock, _ = strconv.Atoi(Stock)
	productFinal.ShipsFrom = ShipsFrom
	productFinal.Offers.Seller_id = Get_id_seller(productFinal.UrlProduct)
	return productFinal
}
func ConvertNumericAbbr(s string) int {
	mapping := make(map[string]int)
	mapping[`k`] = 1000
	mapping[`m`] = 1000000
	var ans float32
	re := regexp.MustCompile(`(\d+([\,]*\d+)*)+(k|m)*`)
	num := re.FindAllString(s, -1)[0]
	if strings.Contains(num, `k`) {
		num = strings.ReplaceAll(num, `k`, ``)
		num = strings.ReplaceAll(num, `,`, `.`)
		num, _ := strconv.ParseFloat(num, 2)
		ans = float32(num) * float32(mapping[`k`])
	} else if strings.Contains(num, `m`) {
		num = strings.ReplaceAll(num, `m`, ``)
		num = strings.ReplaceAll(num, `,`, `.`)
		num, _ := strconv.ParseFloat(num, 2)
		ans = float32(num) * float32(mapping[`m`])
	} else {
		num, _ := strconv.ParseFloat(num, 2)
		if strings.Contains(s, `and`) {
			ans = float32(num) + 1
		} else {
			ans = float32(num)
		}
	}
	return int(ans)
}
func Get_id_seller(product_url string) int64 {
	list_token := strings.Split(product_url, ".")
	fmt.Println(product_url)
	u, err := strconv.ParseInt(list_token[2], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return u
}
