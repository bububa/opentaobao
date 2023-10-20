package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionActivityDeleteAPIResponse 删除档期活动 API返回值
// alibaba.price.promotion.activity.delete
//
// 删除盒马帮档期活动
type AlibabaPricePromotionActivityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaPricePromotionActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPricePromotionActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPricePromotionActivityDeleteAPIResponseModel).Reset()
}

// AlibabaPricePromotionActivityDeleteAPIResponseModel is 删除档期活动 成功返回结果
type AlibabaPricePromotionActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaPricePromotionActivityDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPricePromotionActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPricePromotionActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPricePromotionActivityDeleteAPIResponse)
	},
}

// GetAlibabaPricePromotionActivityDeleteAPIResponse 从 sync.Pool 获取 AlibabaPricePromotionActivityDeleteAPIResponse
func GetAlibabaPricePromotionActivityDeleteAPIResponse() *AlibabaPricePromotionActivityDeleteAPIResponse {
	return poolAlibabaPricePromotionActivityDeleteAPIResponse.Get().(*AlibabaPricePromotionActivityDeleteAPIResponse)
}

// ReleaseAlibabaPricePromotionActivityDeleteAPIResponse 将 AlibabaPricePromotionActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPricePromotionActivityDeleteAPIResponse(v *AlibabaPricePromotionActivityDeleteAPIResponse) {
	v.Reset()
	poolAlibabaPricePromotionActivityDeleteAPIResponse.Put(v)
}
