package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardOpenAPIResponse 标准开卡流程 API返回值
// alibaba.alsc.crm.card.open
//
// 标准开卡流程
type AlibabaAlscCrmCardOpenAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardOpenAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardOpenAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardOpenAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardOpenAPIResponseModel is 标准开卡流程 成功返回结果
type AlibabaAlscCrmCardOpenAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_open_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardOpenAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardOpenAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardOpenAPIResponse)
	},
}

// GetAlibabaAlscCrmCardOpenAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardOpenAPIResponse
func GetAlibabaAlscCrmCardOpenAPIResponse() *AlibabaAlscCrmCardOpenAPIResponse {
	return poolAlibabaAlscCrmCardOpenAPIResponse.Get().(*AlibabaAlscCrmCardOpenAPIResponse)
}

// ReleaseAlibabaAlscCrmCardOpenAPIResponse 将 AlibabaAlscCrmCardOpenAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardOpenAPIResponse(v *AlibabaAlscCrmCardOpenAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardOpenAPIResponse.Put(v)
}
