package fpm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaspopenpaymentsyncAPIResponse 付款单同步 API返回值
// alibaba.sp.open.payment.sync
//
// 新康众弹外同步付款数据
type AlibabaspopenpaymentsyncAPIResponse struct {
	model.CommonResponse
	AlibabaspopenpaymentsyncAPIResponseModel
}

// AlibabaspopenpaymentsyncAPIResponseModel is 付款单同步 成功返回结果
type AlibabaspopenpaymentsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_sp_open_payment_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果处理消息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 是否处理成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
