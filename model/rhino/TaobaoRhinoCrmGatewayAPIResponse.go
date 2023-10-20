package rhino

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoCrmGatewayAPIResponse crm实体变更回调接口 API返回值
// taobao.rhino.crm.gateway
//
// crm实体变更回调接口
type TaobaoRhinoCrmGatewayAPIResponse struct {
	model.CommonResponse
	TaobaoRhinoCrmGatewayAPIResponseModel
}

// TaobaoRhinoCrmGatewayAPIResponseModel is crm实体变更回调接口 成功返回结果
type TaobaoRhinoCrmGatewayAPIResponseModel struct {
	XMLName xml.Name `xml:"rhino_crm_gateway_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息，如果失败时，错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码，200成功，500失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 是否成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
