package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomwttopentradecreateorderAPIResponse 充值送活动下单接口 API返回值
// alibaba.alicom.wtt.opentrade.createorder
//
// 提供给话费宝创建淘宝订单
type AlibabaalicomwttopentradecreateorderAPIResponse struct {
	model.CommonResponse
	AlibabaalicomwttopentradecreateorderAPIResponseModel
}

// AlibabaalicomwttopentradecreateorderAPIResponseModel is 充值送活动下单接口 成功返回结果
type AlibabaalicomwttopentradecreateorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_wtt_opentrade_createorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
