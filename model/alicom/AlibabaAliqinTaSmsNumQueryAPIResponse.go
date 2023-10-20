package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTaSmsNumQueryAPIResponse 短信查询 API返回值
// alibaba.aliqin.ta.sms.num.query
//
// 查询短信发送揭露
type AlibabaAliqinTaSmsNumQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinTaSmsNumQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinTaSmsNumQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinTaSmsNumQueryAPIResponseModel).Reset()
}

// AlibabaAliqinTaSmsNumQueryAPIResponseModel is 短信查询 成功返回结果
type AlibabaAliqinTaSmsNumQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_ta_sms_num_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Values []FcPartnerSmsDetailDto `json:"values,omitempty" xml:"values>fc_partner_sms_detail_dto,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinTaSmsNumQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Values = m.Values[:0]
	m.CurrentPage = 0
	m.PageSize = 0
	m.TotalCount = 0
	m.TotalPage = 0
}

var poolAlibabaAliqinTaSmsNumQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaSmsNumQueryAPIResponse)
	},
}

// GetAlibabaAliqinTaSmsNumQueryAPIResponse 从 sync.Pool 获取 AlibabaAliqinTaSmsNumQueryAPIResponse
func GetAlibabaAliqinTaSmsNumQueryAPIResponse() *AlibabaAliqinTaSmsNumQueryAPIResponse {
	return poolAlibabaAliqinTaSmsNumQueryAPIResponse.Get().(*AlibabaAliqinTaSmsNumQueryAPIResponse)
}

// ReleaseAlibabaAliqinTaSmsNumQueryAPIResponse 将 AlibabaAliqinTaSmsNumQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinTaSmsNumQueryAPIResponse(v *AlibabaAliqinTaSmsNumQueryAPIResponse) {
	v.Reset()
	poolAlibabaAliqinTaSmsNumQueryAPIResponse.Put(v)
}
