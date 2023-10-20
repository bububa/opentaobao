package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolStairAdditemAPIResponse 商品池阶梯商品添加 API返回值
// alibaba.wdk.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
type AlibabaWdkMarketingItempoolStairAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolStairAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolStairAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolStairAdditemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolStairAdditemAPIResponseModel is 商品池阶梯商品添加 成功返回结果
type AlibabaWdkMarketingItempoolStairAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_stair_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolStairAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolStairAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolStairAdditemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolStairAdditemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolStairAdditemAPIResponse
func GetAlibabaWdkMarketingItempoolStairAdditemAPIResponse() *AlibabaWdkMarketingItempoolStairAdditemAPIResponse {
	return poolAlibabaWdkMarketingItempoolStairAdditemAPIResponse.Get().(*AlibabaWdkMarketingItempoolStairAdditemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolStairAdditemAPIResponse 将 AlibabaWdkMarketingItempoolStairAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolStairAdditemAPIResponse(v *AlibabaWdkMarketingItempoolStairAdditemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolStairAdditemAPIResponse.Put(v)
}
