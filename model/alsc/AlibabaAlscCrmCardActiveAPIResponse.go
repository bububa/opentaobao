package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardActiveAPIResponse 标准激活卡 API返回值
// alibaba.alsc.crm.card.active
//
// 激活卡
type AlibabaAlscCrmCardActiveAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardActiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardActiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardActiveAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardActiveAPIResponseModel is 标准激活卡 成功返回结果
type AlibabaAlscCrmCardActiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardActiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardActiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardActiveAPIResponse)
	},
}

// GetAlibabaAlscCrmCardActiveAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardActiveAPIResponse
func GetAlibabaAlscCrmCardActiveAPIResponse() *AlibabaAlscCrmCardActiveAPIResponse {
	return poolAlibabaAlscCrmCardActiveAPIResponse.Get().(*AlibabaAlscCrmCardActiveAPIResponse)
}

// ReleaseAlibabaAlscCrmCardActiveAPIResponse 将 AlibabaAlscCrmCardActiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardActiveAPIResponse(v *AlibabaAlscCrmCardActiveAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardActiveAPIResponse.Put(v)
}
