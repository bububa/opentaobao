package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountDaycostGetAPIResponse 查询今日消耗 API返回值
// alibaba.scbp.account.daycost.get
//
// 查询今日消耗
type AlibabaScbpAccountDaycostGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAccountDaycostGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAccountDaycostGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAccountDaycostGetAPIResponseModel).Reset()
}

// AlibabaScbpAccountDaycostGetAPIResponseModel is 查询今日消耗 成功返回结果
type AlibabaScbpAccountDaycostGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_account_daycost_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回今日消耗，单位元，两位小数
	DayCost string `json:"day_cost,omitempty" xml:"day_cost,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAccountDaycostGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DayCost = ""
}

var poolAlibabaScbpAccountDaycostGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAccountDaycostGetAPIResponse)
	},
}

// GetAlibabaScbpAccountDaycostGetAPIResponse 从 sync.Pool 获取 AlibabaScbpAccountDaycostGetAPIResponse
func GetAlibabaScbpAccountDaycostGetAPIResponse() *AlibabaScbpAccountDaycostGetAPIResponse {
	return poolAlibabaScbpAccountDaycostGetAPIResponse.Get().(*AlibabaScbpAccountDaycostGetAPIResponse)
}

// ReleaseAlibabaScbpAccountDaycostGetAPIResponse 将 AlibabaScbpAccountDaycostGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAccountDaycostGetAPIResponse(v *AlibabaScbpAccountDaycostGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAccountDaycostGetAPIResponse.Put(v)
}
