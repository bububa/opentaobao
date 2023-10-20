package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountJourQueryInfoAPIResponse 查询账户流水信息 API返回值
// alibaba.fundplatform.account.jour.query.info
//
// 外部查询账户流水信息
type AlibabaFundplatformAccountJourQueryInfoAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformAccountJourQueryInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountJourQueryInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformAccountJourQueryInfoAPIResponseModel).Reset()
}

// AlibabaFundplatformAccountJourQueryInfoAPIResponseModel is 查询账户流水信息 成功返回结果
type AlibabaFundplatformAccountJourQueryInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_account_jour_query_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountJourQueryInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFundplatformAccountJourQueryInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformAccountJourQueryInfoAPIResponse)
	},
}

// GetAlibabaFundplatformAccountJourQueryInfoAPIResponse 从 sync.Pool 获取 AlibabaFundplatformAccountJourQueryInfoAPIResponse
func GetAlibabaFundplatformAccountJourQueryInfoAPIResponse() *AlibabaFundplatformAccountJourQueryInfoAPIResponse {
	return poolAlibabaFundplatformAccountJourQueryInfoAPIResponse.Get().(*AlibabaFundplatformAccountJourQueryInfoAPIResponse)
}

// ReleaseAlibabaFundplatformAccountJourQueryInfoAPIResponse 将 AlibabaFundplatformAccountJourQueryInfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformAccountJourQueryInfoAPIResponse(v *AlibabaFundplatformAccountJourQueryInfoAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformAccountJourQueryInfoAPIResponse.Put(v)
}
