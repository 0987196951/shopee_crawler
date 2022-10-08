package main

import (
	"fmt"

	//	"shopee.rd/crawler"
	//	"shopee.rd/database"
	"shopee.rd/utils"
)

func main() {
	notifier := utils.Get_notifier()
	fmt.Println(notifier.SendMessage("hello I'm Bot"))
}

// databaseCnt := database.Connect_to_database(utils.URI_DATABASE)
// product := crawler.GetProduct("https://shopee.vn/ÁO-KHOÁC-KAKI-NAM-FORM-SIÊU-RỘNG-80KG-i.31605840.5134451160")
// fmt.Println(product.ProductID)
// err := database.Delete_product_by_id(databaseCnt.Collection("product"), product.ProductID)
// fmt.Println(err)
// //fmt.Println(product.UrlProduct)
// database.Disconnect_database(databaseCnt)

// func Get_total_product() int {
// 	list_link_category := crawler.GetTreeCategory()
// 	total_product := 0
// 	for _, link := range list_link_category {
// 		fmt.Println(link)
// 		data := crawler.Get_data_from_api_filter(link)
// 		total_product += int(data.Facets[0].Category.Parent_category_detail.Count)
// 	}
// 	return total_product
// }
