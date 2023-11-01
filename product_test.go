package vend

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestListProducts(t *testing.T) {

	client := NewClient(
		os.Getenv("DOMAIN_PREFIX"),
		os.Getenv("API_VERSION"),
		os.Getenv("TOKEN"),
	)
	pageSize := 1000
	deleted := true
	resp, err := client.Product.List(ProductListParams{
		PageSize: &pageSize,
		Deleted:  &deleted,
	})
	if err != nil {
		t.Errorf("Tag.List returned error: %v", err)
	}

	bolB, _ := json.Marshal(len(resp.Data))
	fmt.Println(string(bolB))

}

func TestSearchProducts(t *testing.T) {

	client := NewClient(
		os.Getenv("DOMAIN_PREFIX"),
		os.Getenv("API_VERSION"),
		os.Getenv("TOKEN"),
	)
	skus := []string{"J9-VNC6-KNWZ", "20001055464", "STR_10001024557", "20001056432", "HH-20171026-038", "20001086712", "20001083855", "20001120448", "20001035587", "20001060606", "20001051721", "20001087052", "20001081107", "Classic Baseball Arch_00012", "20001082251", "RPGClass-SS-020", "RPGClass-SS-010", "RPGClass-SS-004", "20001059522", "HH20201030-004", "20001112576", "20180129A-036", "20001089776", "20001092482", "20001089463", "20001064408", "20001058951", "V047-7w-WHT", "20001046986", "20001041185", "20001102309", "20001088376", "20001063699", "CSH-269", "20001089804", "20001079251", "20001118057", "SOLS-0060", "20001089804", "20001042980", "20001051709", "CSLS-211", "20001059524", "20001059544", "STR_20001045939", "20001056962", "20001108407", "20001113928", "20001063812", "20001054627", "20180213A-181", "RPGClass-SS-041", "20001107380", "20001073797", "20001063700", "20001063701", "20001038189", "20001038790", "20001063758", "CSLA-RS4400-022", "CSLA-RS4400-021", "20001078496", "CSH-152", "J9-VNC6-KNWZ", "20001079250", "SC-RQO0-YKJ8", "20001051754", "CSLA-039", "20001083874"}
	resp, err := client.Product.Search(ProductSearchParams{
		SKUs: &skus,
	})
	if err != nil {
		t.Errorf("Tag.List returned error: %v", err)
	}

	bolB, _ := json.Marshal(len(resp.Data))
	fmt.Println(string(bolB))

	// Loop and print sku and categories
	for _, product := range resp.Data {
		fmt.Println(*product.Sku)
		if product.Type != nil {
			fmt.Println(*product.Type.Name)
		}
		if product.Categories != nil {
			for _, category := range *product.Categories {
				fmt.Println(*category.Name)
			}
		}
	}

}

func TestGetProduct(t *testing.T) {

	client := NewClient(
		os.Getenv("DOMAIN_PREFIX"),
		os.Getenv("API_VERSION"),
		os.Getenv("TOKEN"),
	)

	resp, err := client.Product.Get(ProductGetParams{
		ProductID: "b8ca3a6e-7206-11e4-efc6-9775a40a1909",
	})
	if err != nil {
		t.Errorf("Tag.List returned error: %v", err)
	}

	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}
