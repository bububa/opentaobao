package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectAccountListAPIResponse 账户-报表 API返回值
// alibaba.scbp.effect.account.list
//
// 账户-报表,支持最近7天，最近30天，以及180天内时间区间。
type AlibabaScbpEffectAccountListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpEffectAccountListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpEffectAccountListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpEffectAccountListAPIResponseModel).Reset()
}

// AlibabaScbpEffectAccountListAPIResponseModel is 账户-报表 成功返回结果
type AlibabaScbpEffectAccountListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_effect_account_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账户效果数据列表
	AccountReportList []AccountEffectDto `json:"account_report_list,omitempty" xml:"account_report_list>account_effect_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpEffectAccountListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AccountReportList = m.AccountReportList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpEffectAccountListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpEffectAccountListAPIResponse)
	},
}

// GetAlibabaScbpEffectAccountListAPIResponse 从 sync.Pool 获取 AlibabaScbpEffectAccountListAPIResponse
func GetAlibabaScbpEffectAccountListAPIResponse() *AlibabaScbpEffectAccountListAPIResponse {
	return poolAlibabaScbpEffectAccountListAPIResponse.Get().(*AlibabaScbpEffectAccountListAPIResponse)
}

// ReleaseAlibabaScbpEffectAccountListAPIResponse 将 AlibabaScbpEffectAccountListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpEffectAccountListAPIResponse(v *AlibabaScbpEffectAccountListAPIResponse) {
	v.Reset()
	poolAlibabaScbpEffectAccountListAPIResponse.Put(v)
}
