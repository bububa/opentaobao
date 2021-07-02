package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsMessageDirectBatchsendAPIResponse 聚石塔新短信发送接口 API返回值
// taobao.jst.sms.message.direct.batchsend
//
// 聚石塔所见即所得的短信发送接口
type TaobaoJstSmsMessageDirectBatchsendAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsMessageDirectBatchsendAPIResponseModel
}

// TaobaoJstSmsMessageDirectBatchsendAPIResponseModel is 聚石塔新短信发送接口 成功返回结果
type TaobaoJstSmsMessageDirectBatchsendAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_message_direct_batchsend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短信回执码
	Module string `json:"module,omitempty" xml:"module,omitempty"`
}
