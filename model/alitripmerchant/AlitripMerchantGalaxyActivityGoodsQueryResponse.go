package alitripmerchant

// AlitripmerchantgalaxyactivitygoodsqueryResponse 结构体
type AlitripmerchantgalaxyactivitygoodsqueryResponse struct {
	// 奖品信息数据
	Contents []ActivityDrawUserGoodsVo `json:"contents,omitempty" xml:"contents>activity_draw_user_goods_vo,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
