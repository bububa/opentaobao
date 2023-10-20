package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse 储值核销预先校验 API返回值
// alibaba.alsc.crm.recharge.dedutprecheck.get
//
// 储值核销预先校验接口
type AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmRechargeDedutprecheckGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmRechargeDedutprecheckGetAPIResponseModel).Reset()
}

// AlibabaAlscCrmRechargeDedutprecheckGetAPIResponseModel is 储值核销预先校验 成功返回结果
type AlibabaAlscCrmRechargeDedutprecheckGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_recharge_dedutprecheck_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmRechargeDedutprecheckGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmRechargeDedutprecheckGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse)
	},
}

// GetAlibabaAlscCrmRechargeDedutprecheckGetAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse
func GetAlibabaAlscCrmRechargeDedutprecheckGetAPIResponse() *AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse {
	return poolAlibabaAlscCrmRechargeDedutprecheckGetAPIResponse.Get().(*AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse)
}

// ReleaseAlibabaAlscCrmRechargeDedutprecheckGetAPIResponse 将 AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmRechargeDedutprecheckGetAPIResponse(v *AlibabaAlscCrmRechargeDedutprecheckGetAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmRechargeDedutprecheckGetAPIResponse.Put(v)
}
