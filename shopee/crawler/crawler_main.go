package crawler

func Crawler_main() {
	categories := GetTreeCategory()
	var link_products []string
	for _, category := range categories {
		link_products = append(link_products, Get_list_link_product_from_category(category)...)
	}
	var products []Product
	for _, link_product := range link_products {
		products = append(products, GetProduct(link_product))
	}
}
