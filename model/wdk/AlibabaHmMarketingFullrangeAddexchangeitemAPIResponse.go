package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse 全场增加换购品 API返回值
// alibaba.hm.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeAddexchangeitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingFullrangeAddexchangeitemAPIResponseModel).Reset()
}

// AlibabaHmMarketingFullrangeAddexchangeitemAPIResponseModel is 全场增加换购品 成功返回结果
type AlibabaHmMarketingFullrangeAddexchangeitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_addexchangeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeAddexchangeitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingFullrangeAddexchangeitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse)
	},
}

// GetAlibabaHmMarketingFullrangeAddexchangeitemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse
func GetAlibabaHmMarketingFullrangeAddexchangeitemAPIResponse() *AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse {
	return poolAlibabaHmMarketingFullrangeAddexchangeitemAPIResponse.Get().(*AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse)
}

// ReleaseAlibabaHmMarketingFullrangeAddexchangeitemAPIResponse 将 AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeAddexchangeitemAPIResponse(v *AlibabaHmMarketingFullrangeAddexchangeitemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeAddexchangeitemAPIResponse.Put(v)
}
