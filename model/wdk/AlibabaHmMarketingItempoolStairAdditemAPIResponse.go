package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolStairAdditemAPIResponse 商品池阶梯商品添加 API返回值
// alibaba.hm.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
type AlibabaHmMarketingItempoolStairAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolStairAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolStairAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolStairAdditemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolStairAdditemAPIResponseModel is 商品池阶梯商品添加 成功返回结果
type AlibabaHmMarketingItempoolStairAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_stair_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolStairAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolStairAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolStairAdditemAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolStairAdditemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolStairAdditemAPIResponse
func GetAlibabaHmMarketingItempoolStairAdditemAPIResponse() *AlibabaHmMarketingItempoolStairAdditemAPIResponse {
	return poolAlibabaHmMarketingItempoolStairAdditemAPIResponse.Get().(*AlibabaHmMarketingItempoolStairAdditemAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolStairAdditemAPIResponse 将 AlibabaHmMarketingItempoolStairAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolStairAdditemAPIResponse(v *AlibabaHmMarketingItempoolStairAdditemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolStairAdditemAPIResponse.Put(v)
}
