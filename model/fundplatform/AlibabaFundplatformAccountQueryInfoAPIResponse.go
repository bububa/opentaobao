package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountQueryInfoAPIResponse 查询账户信息 API返回值
// alibaba.fundplatform.account.query.info
//
// 外部查询资金平台用户账户信息
type AlibabaFundplatformAccountQueryInfoAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformAccountQueryInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountQueryInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformAccountQueryInfoAPIResponseModel).Reset()
}

// AlibabaFundplatformAccountQueryInfoAPIResponseModel is 查询账户信息 成功返回结果
type AlibabaFundplatformAccountQueryInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_account_query_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountQueryInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFundplatformAccountQueryInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformAccountQueryInfoAPIResponse)
	},
}

// GetAlibabaFundplatformAccountQueryInfoAPIResponse 从 sync.Pool 获取 AlibabaFundplatformAccountQueryInfoAPIResponse
func GetAlibabaFundplatformAccountQueryInfoAPIResponse() *AlibabaFundplatformAccountQueryInfoAPIResponse {
	return poolAlibabaFundplatformAccountQueryInfoAPIResponse.Get().(*AlibabaFundplatformAccountQueryInfoAPIResponse)
}

// ReleaseAlibabaFundplatformAccountQueryInfoAPIResponse 将 AlibabaFundplatformAccountQueryInfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformAccountQueryInfoAPIResponse(v *AlibabaFundplatformAccountQueryInfoAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformAccountQueryInfoAPIResponse.Put(v)
}
