package icbudropshipping

// DistributionSaleProductRequest 
type DistributionSaleProductRequest struct {
    // productId List，max size is 10
    ProductIds   []int64 `json:"product_ids,omitempty" xml:"product_ids>int64,omitempty"`
}
