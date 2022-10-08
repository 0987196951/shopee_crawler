package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type API_Filter_Brand struct {
	Data Data_Filter `json:"data"`
}
type Data_Filter struct {
	Filter_configuration Filter_configuration `json:"filter_configuration"`
}
type Filter_configuration struct {
	Dynamic_Filter_Group_Data Dynamic_Filter_Group_Data `json:"dynamic_filter_group_data"`
}
type Dynamic_Filter_Group_Data struct {
	Location []Location `json:"locations"`
	Brands   []Brand    `json:"brands"`
	Facets   []Facet    `json:"facets"`
}
type Location struct {
	Name         string  `json:"name"`
	Display_name string  `json:"display_name"`
	Tag_ids      []int64 `json:"tag_ids"`
}
type Brand struct {
	Name     string `json:"name"`
	Brand_id int64  `json:"brandid"`
}
type Facet struct {
	Category Category_In_Facet `json:"category"`
	Catid    int64             `json:"catid"`
	Count    int64             `json:"count"`
}
type Category_In_Facet struct {
	Parent_ids             []int64 `json:"parentids"`
	Display_name           string  `json:"display_name"`
	Parent_category_detail *Facet  `json:"parent_category_detail"`
}

func Get_data_from_api_filter(link string) Dynamic_Filter_Group_Data {
	response, err := http.Get(link)
	if response.Status != "200 OK" && err != nil {
		fmt.Println(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var api API_Filter_Brand
	json.Unmarshal(responseData, &api)
	return api.Data.Filter_configuration.Dynamic_Filter_Group_Data
}
