package rhino

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoCrmReviewDeliveryAPIResponse crm实体预询期 API返回值
// taobao.rhino.crm.review.delivery
//
// crm实体预询期
type TaobaoRhinoCrmReviewDeliveryAPIResponse struct {
	model.CommonResponse
	TaobaoRhinoCrmReviewDeliveryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRhinoCrmReviewDeliveryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRhinoCrmReviewDeliveryAPIResponseModel).Reset()
}

// TaobaoRhinoCrmReviewDeliveryAPIResponseModel is crm实体预询期 成功返回结果
type TaobaoRhinoCrmReviewDeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"rhino_crm_review_delivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息，错误时展示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码，成功200/失败500
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 是否成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRhinoCrmReviewDeliveryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.RetCode = 0
	m.Content = false
}

var poolTaobaoRhinoCrmReviewDeliveryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRhinoCrmReviewDeliveryAPIResponse)
	},
}

// GetTaobaoRhinoCrmReviewDeliveryAPIResponse 从 sync.Pool 获取 TaobaoRhinoCrmReviewDeliveryAPIResponse
func GetTaobaoRhinoCrmReviewDeliveryAPIResponse() *TaobaoRhinoCrmReviewDeliveryAPIResponse {
	return poolTaobaoRhinoCrmReviewDeliveryAPIResponse.Get().(*TaobaoRhinoCrmReviewDeliveryAPIResponse)
}

// ReleaseTaobaoRhinoCrmReviewDeliveryAPIResponse 将 TaobaoRhinoCrmReviewDeliveryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRhinoCrmReviewDeliveryAPIResponse(v *TaobaoRhinoCrmReviewDeliveryAPIResponse) {
	v.Reset()
	poolTaobaoRhinoCrmReviewDeliveryAPIResponse.Put(v)
}
