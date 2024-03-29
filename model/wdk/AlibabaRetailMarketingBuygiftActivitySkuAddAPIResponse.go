package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse 添加单品买赠活动商品 API返回值
// alibaba.retail.marketing.buygift.activity.sku.add
//
// 新增或更新单品买赠活动商品信息【同城零售】
type AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponseModel is 添加单品买赠活动商品 成功返回结果
type AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse
func GetAlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse() *AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse {
	return poolAlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse.Get().(*AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse 将 AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse(v *AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse.Put(v)
}
