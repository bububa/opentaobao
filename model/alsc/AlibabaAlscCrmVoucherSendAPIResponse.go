package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmVoucherSendAPIResponse 发送券给指定用户 API返回值
// alibaba.alsc.crm.voucher.send
//
// 发送券给指定用户
type AlibabaAlscCrmVoucherSendAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmVoucherSendAPIResponseModel
}

// AlibabaAlscCrmVoucherSendAPIResponseModel is 发送券给指定用户 成功返回结果
type AlibabaAlscCrmVoucherSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_voucher_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
