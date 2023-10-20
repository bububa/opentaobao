package alicom

import (
	"encoding/xml"

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
