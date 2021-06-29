package ascpqcc

// UpdateSampleData 
type UpdateSampleData struct {
    // 厂商商品尺寸
    VendorSize   string `json:"vendor_size,omitempty" xml:"vendor_size,omitempty"`
    // 业务系统商品编号
    BizItemId   string `json:"biz_item_id,omitempty" xml:"biz_item_id,omitempty"`
}
