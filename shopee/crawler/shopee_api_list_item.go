package crawler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"shopee.rd/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var const_locations = []string{"Hà Nội", "TP. Hồ Chí Minh", "Thái Nguyên", "Vĩnh Phúc",
	"Hải Phòng", "Đồng Nai", "Hưng Yên", "Bình Dương", "Bắc Ninh", "Đà Nẵng", "Quảng Ninh",
	"Hải Dương", "Nam Định", "Cần Thơ", "Phú Thọ", "Bà Rịa - Vũng Tàu", "Đắk Lắk", "Thanh Hóa", "Thái Bình", "Nước ngoài"}

type ProductF struct {
	Type string `json:"@type"`
	Url  string `json:"url"`
}

func Get_list_link_product_from_category(category Category) []string {
	link_filter := Create_link_filter(category.Catid)
	api_filter := Get_data_from_api_filter(link_filter)
	fmt.Println(link_filter)
	brands := api_filter.Brands
	facets := api_filter.Facets
	var list_link_products []string
	for _, location := range const_locations {
		for _, facet := range facets {
			link := Create_link_from_name_and_id(category)
			if location == "Hà Nội" || location == "TP. Hồ Chí Minh" {
				for _, brand := range brands {
					locs := []string{"Hà Nội"}
					list_link_products = append(list_link_products, Get_list_link_product(link, brand.Brand_id, facet.Category.Display_name, locs, "DESC")...)
					locs = []string{"TP. Hồ Chí Minh"}
					list_link_products = append(list_link_products, Get_list_link_product(link, brand.Brand_id, facet.Category.Display_name, locs, "DESC")...)
				}
				list_link_products = append(list_link_products, Get_link_product_with_2_orderby(link, 0, facet.Category.Display_name, []string{location})...)
			} else {
				locs := const_locations[2:]
				list_link_products = append(list_link_products, Get_list_link_product(link, 0, facet.Category.Display_name, locs, "DESC")...)
				break
			}
		}
	}
	return list_link_products
}
func Get_link_product_with_2_orderby(link string, brand int64, facet string, locs []string) []string {
	var list_link_products []string
	list_link_products = append(list_link_products, Get_list_link_product(link, brand, facet, locs, "DESC")...)
	list_link_products = append(list_link_products, Get_list_link_product(link, brand, facet, locs, "INCR")...)
	return list_link_products
}
func Get_list_link_product(link string, brand int64, facet string, locations []string, sort_by string) []string {
	ur := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(ur).MustConnect()
	page := browser.MustPage()
	var links []string
	num_page := 0
	location_str := strings.Join(locations, "&")
	preprocess_link := fmt.Sprintf("%s?brands=%d&facet=%s&locations=%s&sortBy=%s", link, brand, facet, location_str, sort_by)
	limit_page := 1
	for num_page < limit_page {
		link_access := fmt.Sprintf("%s&page=%d", preprocess_link, num_page)
		link_access = strings.ReplaceAll(link_access, " ", "%20")
		fmt.Println(link_access)
		page.MustNavigate(link_access)
		if utils.Is_link_Status_200(link_access) == false || num_page >= 49 {
			break
		}
		time.Sleep(5 * time.Second)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(page.MustHTML()))
		if err != nil {
			fmt.Println("co loi")
		}
		if num_page == 0 {
			limit_page, err = strconv.Atoi(doc.Find("span.shopee-mini-page-controller__total").Text())
		}
		if err != nil {
			fmt.Println("Sai chuyen doi")
			break
		}
		var products []ProductF
		doc.Find("script").Each(func(i int, s *goquery.Selection) {
			attr, exist := s.Attr("type")
			if attr == "application/ld+json" && exist == true {
				var product ProductF
				json.Unmarshal([]byte(s.Text()), &product)
				if product.Type == "Product" {
					products = append(products, product)
				}
			}
		})

		for _, product := range products {
			links = append(links, product.Url)
		}
		num_page++
		fmt.Println(len(links))
	}
	fmt.Println(links[0])
	return links
}
func Create_link_from_name_and_id(category Category) string {
	fmt.Println(category)
	base := "https://shopee.vn/"
	name := strings.ReplaceAll(category.Display_name, " ", "-")
	return fmt.Sprintf("%s%s-cat.%d.%d", base, name, category.Parent_catid, category.Catid)
}
func Create_link_filter(catid int64) string {
	return fmt.Sprintf("https://shopee.vn/api/v4/search/search_filter_config?match_id=%d&page_type=search&scenario=PAGE_CATEGORY", catid)
}

// func get_list_links_from_link_product( url string) []string{

//		page := browser.MustNavigate(url)
//		if Is_link_Status_200(url) == false {
//			return []string{}
//		}
//		time.Sleep(5 * time.Second)
//		doc, err := goquery.NewDocumentFromReader(strings.NewReader(page.MustHTML()))
//		if err != nil {
//			fmt.Println("co loi")
//		}
//		var products []Product
//		doc.Find("script").Each(func(i int, s *goquery.Selection) {
//			attr, exist := s.Attr("type")
//			if attr == "application/ld+json" && exist == true {
//				var product Product
//				json.Unmarshal([]byte(s.Text()), &product)
//				if product.Type == "Product" {
//					products = append(products, product)
//				}
//			}
//		})
//		var links []string
//		for _, product := range products {
//			links = append(links, product.Url)
//		}
//		return links
//	}
// func Is_link_Status_200(link string) bool {
// 	response, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(1)
// 		return false
// 	}
// 	if strings.Compare(response.Status, "200 OK") != 0 {
// 		return false
// 	}
// 	return true
// }

// func Get_links_is_not_exist(links []string) []int {
// 	var index []int
// 	for i, link := range links {
// 		if Is_link_Status_200(link) == false {
// 			index = append(index, i)
// 		}else {
// 			fmt.Println(link)
// 		}
// 	}
// 	return index
// }
