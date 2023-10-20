package viapi

// AliyunViapiGoodstechClassifygoodsData 结构体
type AliyunViapiGoodstechClassifygoodsData struct {
	// 类目预测列表
	CategoryList []Category `json:"category_list,omitempty" xml:"category_list>category,omitempty"`
}
