package icbudropshipping

// DistributionSaleProduct 
type DistributionSaleProduct struct {
    // Ladder delivery List
    LadderPeriodList   []LadderPeriod `json:"ladder_period_list,omitempty" xml:"ladder_period_list>ladder_period,omitempty"`
    // main image url
    MainImageUrl   string `json:"main_image_url,omitempty" xml:"main_image_url,omitempty"`
    // product name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // price Range
    PriceRange   string `json:"price_range,omitempty" xml:"price_range,omitempty"`
    // product id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // product sku list
    ProductSkuList   []ProductSku `json:"product_sku_list,omitempty" xml:"product_sku_list>product_sku,omitempty"`
    // product detail Url
    DetailUrl   string `json:"detail_url,omitempty" xml:"detail_url,omitempty"`
    // min order quantity and Price
    MoqAndPrice   *MoqAndPrice `json:"moq_and_price,omitempty" xml:"moq_and_price,omitempty"`
    // product description html，<br /> It is the transferred string, the applicable party needs to reverse the string
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // Determine whether this product can be ordered
    IsCanPlaceOrder   bool `json:"is_can_place_order,omitempty" xml:"is_can_place_order,omitempty"`
    // product keywords
    Keywords   []string `json:"keywords,omitempty" xml:"keywords>string,omitempty"`
    // Store Id
    ECompanyId   string `json:"e_company_id,omitempty" xml:"e_company_id,omitempty"`
    // the top category name
    TopCategoryName   string `json:"top_category_name,omitempty" xml:"top_category_name,omitempty"`
    // the leaf category name
    LeafCategoryName   string `json:"leaf_category_name,omitempty" xml:"leaf_category_name,omitempty"`
    // product image list
    ImageUrlList   []string `json:"image_url_list,omitempty" xml:"image_url_list>string,omitempty"`
}
