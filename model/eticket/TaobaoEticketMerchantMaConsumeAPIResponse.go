package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoeticketmerchantmaconsumeAPIResponse 电子凭证核销接口 API返回值
// taobao.eticket.merchant.ma.consume
//
// 电子凭证核销接口
type TaobaoeticketmerchantmaconsumeAPIResponse struct {
	model.CommonResponse
	TaobaoeticketmerchantmaconsumeAPIResponseModel
}

// TaobaoeticketmerchantmaconsumeAPIResponseModel is 电子凭证核销接口 成功返回结果
type TaobaoeticketmerchantmaconsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 系统自动生成
	RespBody *ConsumeMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}
