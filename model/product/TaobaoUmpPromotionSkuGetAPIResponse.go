package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpPromotionSkuGetAPIResponse 商品优惠详情查询 API返回值
// taobao.ump.promotion.sku.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TaobaoUmpPromotionSkuGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpPromotionSkuGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpPromotionSkuGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpPromotionSkuGetAPIResponseModel).Reset()
}

// TaobaoUmpPromotionSkuGetAPIResponseModel is 商品优惠详情查询 成功返回结果
type TaobaoUmpPromotionSkuGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_promotion_sku_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Models string `json:"models,omitempty" xml:"models,omitempty"`
	// 默认key值
	DefaultModelKey string `json:"default_model_key,omitempty" xml:"default_model_key,omitempty"`
	// 调用失败信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpPromotionSkuGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Models = ""
	m.DefaultModelKey = ""
	m.ErrorMessage = ""
	m.IsSuccess = false
}

var poolTaobaoUmpPromotionSkuGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpPromotionSkuGetAPIResponse)
	},
}

// GetTaobaoUmpPromotionSkuGetAPIResponse 从 sync.Pool 获取 TaobaoUmpPromotionSkuGetAPIResponse
func GetTaobaoUmpPromotionSkuGetAPIResponse() *TaobaoUmpPromotionSkuGetAPIResponse {
	return poolTaobaoUmpPromotionSkuGetAPIResponse.Get().(*TaobaoUmpPromotionSkuGetAPIResponse)
}

// ReleaseTaobaoUmpPromotionSkuGetAPIResponse 将 TaobaoUmpPromotionSkuGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpPromotionSkuGetAPIResponse(v *TaobaoUmpPromotionSkuGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpPromotionSkuGetAPIResponse.Put(v)
}
