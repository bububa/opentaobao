package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBatchActiveAPIResponse 批量激活卡 API返回值
// alibaba.alsc.crm.card.batch.active
//
// 批量激活卡
type AlibabaAlscCrmCardBatchActiveAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardBatchActiveAPIResponseModel
}

// AlibabaAlscCrmCardBatchActiveAPIResponseModel is 批量激活卡 成功返回结果
type AlibabaAlscCrmCardBatchActiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_batch_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
