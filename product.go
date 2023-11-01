package vend

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const productsBasePath = "products"
const searchBasePath = "search"

type ProductServiceOp struct {
	client *Client
}

type ProductService interface {
	List(ProductListParams) (*ProductsResource, error)
	Search(ProductSearchParams) (*ProductsResource, error)
	Get(ProductGetParams) (*ProductResource, error)
}

type ProductListParams struct {
	After    *int64
	Before   *int64
	Deleted  *bool
	PageSize *int
}

type ProductSearchParams struct {
	SKUs             *[]string
	SupplierIDs      *[]string
	BrandIDs         *[]string
	TagIDs           *[]string
	ProductTypeIDs   *[]string
	VariantParentIDs *[]string
}

type ProductGetParams struct {
	ProductID string
}

type ProductResource struct {
	Data Product `json:"data,omitempty"`
}

type ProductsResource struct {
	Data    []Product `json:"data,omitempty"`
	Version Version   `json:"version,omitempty"`
}
type Product struct {
	ID                *string           `json:"id,omitempty"`
	Name              *string           `json:"name,omitempty"`
	VariantName       *string           `json:"variant_name,omitempty"`
	Handle            *string           `json:"handle,omitempty"`
	Sku               *string           `json:"sku,omitempty"`
	SupplierCode      *string           `json:"supplier_code,omitempty"`
	Active            *bool             `json:"active,omitempty"`
	HasInventory      *bool             `json:"has_inventory,omitempty"`
	IsComposite       *bool             `json:"is_composite,omitempty"`
	Description       *string           `json:"description,omitempty"`
	ImageURL          *string           `json:"image_url,omitempty"`
	CreatedAt         *time.Time        `json:"created_at,omitempty"`
	UpdatedAt         *time.Time        `json:"updated_at,omitempty"`
	Source            *string           `json:"source,omitempty"`
	SupplyPrice       *float64          `json:"supply_price,omitempty"`
	Version           *int64            `json:"version,omitempty"`
	Type              *Type             `json:"type,omitempty"`
	Supplier          *Supplier         `json:"supplier,omitempty"`
	Brand             *Brand            `json:"brand,omitempty"`
	VariantOptions    *[]VariantOptions `json:"variant_options,omitempty"`
	Categories        *[]Categories     `json:"categories,omitempty"`
	Images            *[]Images         `json:"images,omitempty"`
	HasVariants       *bool             `json:"has_variants,omitempty"`
	ButtonOrder       *int              `json:"button_order,omitempty"`
	PriceIncludingTax *float64          `json:"price_including_tax,omitempty"`
	PriceExcludingTax *float64          `json:"price_excluding_tax,omitempty"`
	// Attributes        *[]any            `json:"attributes,omitempty"`
	SupplierID        *string   `json:"supplier_id,omitempty"`
	ProductTypeID     *string   `json:"product_type_id,omitempty"`
	BrandID           *string   `json:"brand_id,omitempty"`
	IsActive          *bool     `json:"is_active,omitempty"`
	ImageThumbnailURL *string   `json:"image_thumbnail_url,omitempty"`
	TagIds            *[]string `json:"tag_ids,omitempty"`
}
type Type struct {
	ID      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *int    `json:"version,omitempty"`
}
type Supplier struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Source      *string `json:"source,omitempty"`
	Description *string `json:"description,omitempty"`
	Version     *int    `json:"version,omitempty"`
}
type Brand struct {
	ID      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *int    `json:"version,omitempty"`
}
type VariantOptions struct {
	ID    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
type Categories struct {
	ID      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *int    `json:"version,omitempty"`
}
type Sizes struct {
	Ss       *string `json:"ss,omitempty"`
	Standard *string `json:"standard,omitempty"`
	St       *string `json:"st,omitempty"`
	Original *string `json:"original,omitempty"`
	Thumb    *string `json:"thumb,omitempty"`
	Sl       *string `json:"sl,omitempty"`
	Sm       *string `json:"sm,omitempty"`
}
type Images *struct {
	ID      *string `json:"id,omitempty"`
	URL     *string `json:"url,omitempty"`
	Version *int64  `json:"version,omitempty"`
	Sizes   *Sizes  `json:"sizes,omitempty"`
}
type Version struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

func (s *ProductServiceOp) List(params ProductListParams) (*ProductsResource, error) {

	values := url.Values{}
	if params.After != nil {
		values.Add("after", strconv.FormatInt(*params.After, 10))
	}
	if params.Before != nil {
		values.Add("after", strconv.FormatInt(*params.Before, 10))
	}
	if params.Deleted != nil {
		if *params.Deleted {
			values.Add("deleted", "true")
		} else {
			values.Add("deleted", "false")
		}
	}
	if params.PageSize != nil {
		values.Add("page_size", strconv.Itoa(*params.PageSize))
	}

	path := productsBasePath + "?" + values.Encode()

	var resp ProductsResource

	errRequest := s.client.Request("GET", path, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ProductServiceOp) Search(params ProductSearchParams) (*ProductsResource, error) {

	values := url.Values{}
	values.Add("type", productsBasePath)
	if params.SKUs != nil {
		for _, sku := range *params.SKUs {
			values.Add("sku", strings.ToLower(sku))
		}
	}
	if params.SupplierIDs != nil {
		for _, supplierID := range *params.SupplierIDs {
			values.Add("supplier_id", supplierID)
		}
	}
	if params.BrandIDs != nil {
		for _, brandID := range *params.BrandIDs {
			values.Add("brand_id", brandID)
		}
	}
	if params.TagIDs != nil {
		for _, tagID := range *params.TagIDs {
			values.Add("tag_id", tagID)
		}
	}
	if params.ProductTypeIDs != nil {
		for _, productTypeID := range *params.ProductTypeIDs {
			values.Add("product_type_id", productTypeID)
		}
	}
	if params.VariantParentIDs != nil {
		for _, variantParentID := range *params.VariantParentIDs {
			values.Add("variant_parent_id", variantParentID)
		}
	}

	path := searchBasePath + "?" + values.Encode()

	var resp ProductsResource

	errRequest := s.client.Request("GET", path, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ProductServiceOp) Get(params ProductGetParams) (*ProductResource, error) {

	path := productsBasePath + "/" + params.ProductID

	var resp ProductResource

	errRequest := s.client.Request("GET", path, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
