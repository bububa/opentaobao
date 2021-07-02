package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenOrderBackflowAPIResponse 订单回流接口 API返回值
// alibaba.alsc.crm.open.order.backflow
//
// 回流isv订单接口
type AlibabaAlscCrmOpenOrderBackflowAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmOpenOrderBackflowAPIResponseModel
}

// AlibabaAlscCrmOpenOrderBackflowAPIResponseModel is 订单回流接口 成功返回结果
type AlibabaAlscCrmOpenOrderBackflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_order_backflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
