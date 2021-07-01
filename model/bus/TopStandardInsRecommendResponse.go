package bus

// TopStandardInsRecommendResponse 结构体
type TopStandardInsRecommendResponse struct {
	// 推荐结果
	ResultList []TopInsProduct `json:"result_list,omitempty" xml:"result_list>top_ins_product,omitempty"`
}
