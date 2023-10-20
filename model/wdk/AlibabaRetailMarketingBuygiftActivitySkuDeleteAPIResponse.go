package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse 删除单品买赠活动商品 API返回值
// alibaba.retail.marketing.buygift.activity.sku.delete
//
// 删除单品买赠活动商品信息【同城零售】
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponseModel is 删除单品买赠活动商品 成功返回结果
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse
func GetAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse() *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse {
	return poolAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse.Get().(*AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse 将 AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse(v *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse.Put(v)
}
