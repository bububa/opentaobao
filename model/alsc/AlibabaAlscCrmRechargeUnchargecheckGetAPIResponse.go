package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse 储值账户退充值校验 API返回值
// alibaba.alsc.crm.recharge.unchargecheck.get
//
// 储值账户退充值校验接口
type AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeUnchargecheckGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeUnchargecheckGetAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeUnchargecheckGetAPIResponseModel is 储值账户退充值校验 成功返回结果
type AlibabaAlscCrmRechargeUnchargecheckGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_unchargecheck_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeUnchargecheckGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeUnchargecheckGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeUnchargecheckGetAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse
func GetAlibabaAlscCrmRechargeUnchargecheckGetAPIResponse() *AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse {
	return poolAlibabaAlscCrmRechargeUnchargecheckGetAPIResponse.Get().(*AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeUnchargecheckGetAPIResponse 将 AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeUnchargecheckGetAPIResponse(v *AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeUnchargecheckGetAPIResponse.Put(v)
}
