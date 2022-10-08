package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func Get_links_is_not_exist(links []string) []int {
	var index []int
	for i, link := range links {
		if Is_link_Status_200(link) == false {
			index = append(index, i)
		}
	}
	return index
}
func Is_link_Status_200(link string) bool {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(1)
		return false
	}
	if strings.Compare(response.Status, "200 OK") != 0 {
		return false
	}
	return true
}
