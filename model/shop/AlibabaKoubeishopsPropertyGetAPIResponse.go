package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabakoubeishopspropertygetAPIResponse 口碑店铺列表推荐 API返回值
// alibaba.koubeishops.property.get
//
// 推荐用户附近的美食门店
type AlibabakoubeishopspropertygetAPIResponse struct {
	model.CommonResponse
	AlibabakoubeishopspropertygetAPIResponseModel
}

// AlibabakoubeishopspropertygetAPIResponseModel is 口碑店铺列表推荐 成功返回结果
type AlibabakoubeishopspropertygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_koubeishops_property_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenApiSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
