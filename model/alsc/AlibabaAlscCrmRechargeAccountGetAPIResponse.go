package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeAccountGetAPIResponse 查询储值账户信息 API返回值
// alibaba.alsc.crm.recharge.account.get
//
// 查询储值账户信息接口
type AlibabaAlscCrmRechargeAccountGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeAccountGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeAccountGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeAccountGetAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeAccountGetAPIResponseModel is 查询储值账户信息 成功返回结果
type AlibabaAlscCrmRechargeAccountGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_account_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeAccountGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeAccountGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeAccountGetAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeAccountGetAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeAccountGetAPIResponse
func GetAlibabaAlscCrmRechargeAccountGetAPIResponse() *AlibabaAlscCrmRechargeAccountGetAPIResponse {
	return poolAlibabaAlscCrmRechargeAccountGetAPIResponse.Get().(*AlibabaAlscCrmRechargeAccountGetAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeAccountGetAPIResponse 将 AlibabaAlscCrmRechargeAccountGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeAccountGetAPIResponse(v *AlibabaAlscCrmRechargeAccountGetAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeAccountGetAPIResponse.Put(v)
}
