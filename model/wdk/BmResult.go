package wdk

// BmResult 结构体
type BmResult struct {
	// 结果数据
	DataList []AlibabawdkbmcouponqueryData `json:"data_list,omitempty" xml:"data_list>alibabawdkbmcouponquery_data,omitempty"`
	// 详细结果
	PublishResults []SkuStockPublishResult `json:"publish_results,omitempty" xml:"publish_results>sku_stock_publish_result,omitempty"`
	// 错误码，当失败返回错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 描述信息，当成功返回OK，当失败返回错误描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 额外属性
	ExtData string `json:"ext_data,omitempty" xml:"ext_data,omitempty"`
	// 库存数量
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
