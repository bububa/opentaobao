package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlipaypfmConsumeRecordAPIResponse 记录用户每日消耗卡路里总量 API返回值
// alibaba.alihealth.alipaypfm.consume.record
//
// 记录用户每日消耗卡路里总量
type AlibabaAlihealthAlipaypfmConsumeRecordAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlipaypfmConsumeRecordAPIResponseModel
}

// AlibabaAlihealthAlipaypfmConsumeRecordAPIResponseModel is 记录用户每日消耗卡路里总量 成功返回结果
type AlibabaAlihealthAlipaypfmConsumeRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_consume_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
