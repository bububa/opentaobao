package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneItemExternalRecommendAPIResponse 话费选品能力外放 API返回值
// taobao.phone.item.external.recommend
//
// 话费选品能力外放
type TaobaoPhoneItemExternalRecommendAPIResponse struct {
	model.CommonResponse
	TaobaoPhoneItemExternalRecommendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPhoneItemExternalRecommendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPhoneItemExternalRecommendAPIResponseModel).Reset()
}

// TaobaoPhoneItemExternalRecommendAPIResponseModel is 话费选品能力外放 成功返回结果
type TaobaoPhoneItemExternalRecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"phone_item_external_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data *PhoneRecommendRes `json:"data,omitempty" xml:"data,omitempty"`
	// 响应状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPhoneItemExternalRecommendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizCode = ""
	m.Message = ""
	m.Data = nil
	m.IsSuccess = false
}

var poolTaobaoPhoneItemExternalRecommendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPhoneItemExternalRecommendAPIResponse)
	},
}

// GetTaobaoPhoneItemExternalRecommendAPIResponse 从 sync.Pool 获取 TaobaoPhoneItemExternalRecommendAPIResponse
func GetTaobaoPhoneItemExternalRecommendAPIResponse() *TaobaoPhoneItemExternalRecommendAPIResponse {
	return poolTaobaoPhoneItemExternalRecommendAPIResponse.Get().(*TaobaoPhoneItemExternalRecommendAPIResponse)
}

// ReleaseTaobaoPhoneItemExternalRecommendAPIResponse 将 TaobaoPhoneItemExternalRecommendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPhoneItemExternalRecommendAPIResponse(v *TaobaoPhoneItemExternalRecommendAPIResponse) {
	v.Reset()
	poolTaobaoPhoneItemExternalRecommendAPIResponse.Put(v)
}
