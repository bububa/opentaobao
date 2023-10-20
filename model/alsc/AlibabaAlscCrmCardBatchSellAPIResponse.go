package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBatchSellAPIResponse 批量开卡（售卡） API返回值
// alibaba.alsc.crm.card.batch.sell
//
// 批量开卡（售卡）
type AlibabaAlscCrmCardBatchSellAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardBatchSellAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardBatchSellAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardBatchSellAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardBatchSellAPIResponseModel is 批量开卡（售卡） 成功返回结果
type AlibabaAlscCrmCardBatchSellAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_batch_sell_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardBatchSellAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardBatchSellAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardBatchSellAPIResponse)
	},
}

// GetAlibabaAlscCrmCardBatchSellAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardBatchSellAPIResponse
func GetAlibabaAlscCrmCardBatchSellAPIResponse() *AlibabaAlscCrmCardBatchSellAPIResponse {
	return poolAlibabaAlscCrmCardBatchSellAPIResponse.Get().(*AlibabaAlscCrmCardBatchSellAPIResponse)
}

// ReleaseAlibabaAlscCrmCardBatchSellAPIResponse 将 AlibabaAlscCrmCardBatchSellAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardBatchSellAPIResponse(v *AlibabaAlscCrmCardBatchSellAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardBatchSellAPIResponse.Put(v)
}
