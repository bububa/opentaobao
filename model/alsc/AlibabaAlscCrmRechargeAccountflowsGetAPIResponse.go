package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeAccountflowsGetAPIResponse 分页查询储值流水 API返回值
// alibaba.alsc.crm.recharge.accountflows.get
//
// 增加分页查询储值流水接口
type AlibabaAlscCrmRechargeAccountflowsGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeAccountflowsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeAccountflowsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeAccountflowsGetAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeAccountflowsGetAPIResponseModel is 分页查询储值流水 成功返回结果
type AlibabaAlscCrmRechargeAccountflowsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_accountflows_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeAccountflowsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeAccountflowsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeAccountflowsGetAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeAccountflowsGetAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeAccountflowsGetAPIResponse
func GetAlibabaAlscCrmRechargeAccountflowsGetAPIResponse() *AlibabaAlscCrmRechargeAccountflowsGetAPIResponse {
	return poolAlibabaAlscCrmRechargeAccountflowsGetAPIResponse.Get().(*AlibabaAlscCrmRechargeAccountflowsGetAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeAccountflowsGetAPIResponse 将 AlibabaAlscCrmRechargeAccountflowsGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeAccountflowsGetAPIResponse(v *AlibabaAlscCrmRechargeAccountflowsGetAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeAccountflowsGetAPIResponse.Put(v)
}
