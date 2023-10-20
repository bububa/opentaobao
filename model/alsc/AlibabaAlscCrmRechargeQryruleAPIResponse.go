package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeQryruleAPIResponse 储值规则下行 API返回值
// alibaba.alsc.crm.recharge.qryrule
//
// 储值规则下行
type AlibabaAlscCrmRechargeQryruleAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeQryruleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeQryruleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeQryruleAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeQryruleAPIResponseModel is 储值规则下行 成功返回结果
type AlibabaAlscCrmRechargeQryruleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_qryrule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回模型
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeQryruleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeQryruleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeQryruleAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeQryruleAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeQryruleAPIResponse
func GetAlibabaAlscCrmRechargeQryruleAPIResponse() *AlibabaAlscCrmRechargeQryruleAPIResponse {
	return poolAlibabaAlscCrmRechargeQryruleAPIResponse.Get().(*AlibabaAlscCrmRechargeQryruleAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeQryruleAPIResponse 将 AlibabaAlscCrmRechargeQryruleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeQryruleAPIResponse(v *AlibabaAlscCrmRechargeQryruleAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeQryruleAPIResponse.Put(v)
}
