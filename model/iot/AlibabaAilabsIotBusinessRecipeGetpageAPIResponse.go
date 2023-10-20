package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotbusinessrecipegetpageAPIResponse 分页查询食谱 API返回值
// alibaba.ailabs.iot.business.recipe.getpage
//
// 分页查询食谱数据
type AlibabaailabsiotbusinessrecipegetpageAPIResponse struct {
	model.CommonResponse
	AlibabaailabsiotbusinessrecipegetpageAPIResponseModel
}

// AlibabaailabsiotbusinessrecipegetpageAPIResponseModel is 分页查询食谱 成功返回结果
type AlibabaailabsiotbusinessrecipegetpageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_business_recipe_getpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
