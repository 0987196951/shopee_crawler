package crawler

import (
	"fmt"
	//"github.com/go-rod/rod"
	//"time"
	//"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type API struct {
	Data Data `json:"data"`
}
type Data struct {
	Category_list []Category `json:"category_list"`
}

func GetTreeCategory() []Category {
	response, err := http.Get("https://shopee.vn/api/v4/pages/get_category_tree")
	if err != nil {
		fmt.Println(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var api API
	json.Unmarshal(responseData, &api)
	// for _, link := range list_link_access_category_level_2 {
	// 	fmt.Println(link)
	// }
	var categories []Category
	for _, category := range api.Data.Category_list {
		categories = append(categories, category.Children...)
	}
	return categories
}
func Get_list_category_level_2(api API) []string {
	var list_category_level_2 []string
	list_category_level_1 := api.Data.Category_list
	for _, category_level_1 := range list_category_level_1 {
		//parent_catid := category_level_1.Catid
		for _, category_level_2 := range category_level_1.Children {
			catid := category_level_2.Catid
			//display_name := category_level_2.Display_name
			// link_page := Create_link_category_with_filter(display_name, parent_catid, catid)
			link_page := Create_link_category_with_filter(strconv.FormatInt(catid, 10))
			list_category_level_2 = append(list_category_level_2, link_page)
		}
	}
	return list_category_level_2
}
func Create_link_category_with_filter(catid string) string {
	return fmt.Sprintf("https://shopee.vn/api/v4/search/search_filter_config?match_id=%s&page_type=search&scenario=PAGE_CATEGORY", catid)
}
