package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractOrderCheckuserimeiAPIResponse 金融购机验证设备号 API返回值
// alibaba.interact.order.checkuserimei
//
// 金融购机验证用户身份信息
type AlibabaInteractOrderCheckuserimeiAPIResponse struct {
	model.CommonResponse
	AlibabaInteractOrderCheckuserimeiAPIResponseModel
}

// AlibabaInteractOrderCheckuserimeiAPIResponseModel is 金融购机验证设备号 成功返回结果
type AlibabaInteractOrderCheckuserimeiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_order_checkuserimei_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	SuccessStatus bool `json:"success_status,omitempty" xml:"success_status,omitempty"`
	// 响应数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
