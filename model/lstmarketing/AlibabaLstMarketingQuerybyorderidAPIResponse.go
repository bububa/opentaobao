package lstmarketing

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstMarketingQuerybyorderidAPIResponse 根据订单查询营销信息 API返回值
// alibaba.lst.marketing.querybyorderid
//
// 根据订单查询营销信息
type AlibabaLstMarketingQuerybyorderidAPIResponse struct {
	model.CommonResponse
	AlibabaLstMarketingQuerybyorderidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstMarketingQuerybyorderidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstMarketingQuerybyorderidAPIResponseModel).Reset()
}

// AlibabaLstMarketingQuerybyorderidAPIResponseModel is 根据订单查询营销信息 成功返回结果
type AlibabaLstMarketingQuerybyorderidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_marketing_querybyorderid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaLstMarketingQuerybyorderidResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstMarketingQuerybyorderidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstMarketingQuerybyorderidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstMarketingQuerybyorderidAPIResponse)
	},
}

// GetAlibabaLstMarketingQuerybyorderidAPIResponse 从 sync.Pool 获取 AlibabaLstMarketingQuerybyorderidAPIResponse
func GetAlibabaLstMarketingQuerybyorderidAPIResponse() *AlibabaLstMarketingQuerybyorderidAPIResponse {
	return poolAlibabaLstMarketingQuerybyorderidAPIResponse.Get().(*AlibabaLstMarketingQuerybyorderidAPIResponse)
}

// ReleaseAlibabaLstMarketingQuerybyorderidAPIResponse 将 AlibabaLstMarketingQuerybyorderidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstMarketingQuerybyorderidAPIResponse(v *AlibabaLstMarketingQuerybyorderidAPIResponse) {
	v.Reset()
	poolAlibabaLstMarketingQuerybyorderidAPIResponse.Put(v)
}
