package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse 全场增加换购品 API返回值
// alibaba.wdk.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel is 全场增加换购品 成功返回结果
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_addexchangeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse
func GetAlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse() *AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse {
	return poolAlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse.Get().(*AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse 将 AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse(v *AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse.Put(v)
}
