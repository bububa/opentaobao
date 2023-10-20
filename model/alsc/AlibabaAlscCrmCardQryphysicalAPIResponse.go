package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQryphysicalAPIResponse 查询物理卡 API返回值
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
type AlibabaAlscCrmCardQryphysicalAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardQryphysicalAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardQryphysicalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCardQryphysicalAPIResponseModel).Reset()
}

// AlibabaAlscCrmCardQryphysicalAPIResponseModel is 查询物理卡 成功返回结果
type AlibabaAlscCrmCardQryphysicalAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_qryphysical_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCardQryphysicalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCardQryphysicalAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCardQryphysicalAPIResponse)
	},
}

// GetAlibabaAlscCrmCardQryphysicalAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCardQryphysicalAPIResponse
func GetAlibabaAlscCrmCardQryphysicalAPIResponse() *AlibabaAlscCrmCardQryphysicalAPIResponse {
	return poolAlibabaAlscCrmCardQryphysicalAPIResponse.Get().(*AlibabaAlscCrmCardQryphysicalAPIResponse)
}

// ReleaseAlibabaAlscCrmCardQryphysicalAPIResponse 将 AlibabaAlscCrmCardQryphysicalAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCardQryphysicalAPIResponse(v *AlibabaAlscCrmCardQryphysicalAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCardQryphysicalAPIResponse.Put(v)
}
