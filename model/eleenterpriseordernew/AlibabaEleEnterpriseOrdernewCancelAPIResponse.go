package eleenterpriseordernew

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriseordernewcancelAPIResponse 订单取消 API返回值
// alibaba.ele.enterprise.ordernew.cancel
//
// 订单取消
type AlibabaeleenterpriseordernewcancelAPIResponse struct {
	model.CommonResponse
	AlibabaeleenterpriseordernewcancelAPIResponseModel
}

// AlibabaeleenterpriseordernewcancelAPIResponseModel is 订单取消 成功返回结果
type AlibabaeleenterpriseordernewcancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}
