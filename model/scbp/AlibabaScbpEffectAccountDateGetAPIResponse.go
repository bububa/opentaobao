package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectAccountDateGetAPIResponse 获取最近报表生成时间 API返回值
// alibaba.scbp.effect.account.date.get
//
// 获取最近报表生成时间,格式为yyyy-MM-dd
type AlibabaScbpEffectAccountDateGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpEffectAccountDateGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpEffectAccountDateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpEffectAccountDateGetAPIResponseModel).Reset()
}

// AlibabaScbpEffectAccountDateGetAPIResponseModel is 获取最近报表生成时间 成功返回结果
type AlibabaScbpEffectAccountDateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_effect_account_date_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最近生成报表的时间(US)
	ReportDate string `json:"report_date,omitempty" xml:"report_date,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpEffectAccountDateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReportDate = ""
}

var poolAlibabaScbpEffectAccountDateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpEffectAccountDateGetAPIResponse)
	},
}

// GetAlibabaScbpEffectAccountDateGetAPIResponse 从 sync.Pool 获取 AlibabaScbpEffectAccountDateGetAPIResponse
func GetAlibabaScbpEffectAccountDateGetAPIResponse() *AlibabaScbpEffectAccountDateGetAPIResponse {
	return poolAlibabaScbpEffectAccountDateGetAPIResponse.Get().(*AlibabaScbpEffectAccountDateGetAPIResponse)
}

// ReleaseAlibabaScbpEffectAccountDateGetAPIResponse 将 AlibabaScbpEffectAccountDateGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpEffectAccountDateGetAPIResponse(v *AlibabaScbpEffectAccountDateGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpEffectAccountDateGetAPIResponse.Put(v)
}
