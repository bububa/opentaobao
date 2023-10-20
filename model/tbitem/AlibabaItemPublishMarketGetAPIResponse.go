package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemPublishMarketGetAPIResponse 获取商家可发布商品的市场信息 API返回值
// alibaba.item.publish.market.get
//
// 获取商家可发布商品的市场信息
type AlibabaItemPublishMarketGetAPIResponse struct {
	model.CommonResponse
	AlibabaItemPublishMarketGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemPublishMarketGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemPublishMarketGetAPIResponseModel).Reset()
}

// AlibabaItemPublishMarketGetAPIResponseModel is 获取商家可发布商品的市场信息 成功返回结果
type AlibabaItemPublishMarketGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_publish_market_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商家可发布的市场列表,多个以逗号(,)分隔
	Markets string `json:"markets,omitempty" xml:"markets,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemPublishMarketGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Markets = ""
}

var poolAlibabaItemPublishMarketGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemPublishMarketGetAPIResponse)
	},
}

// GetAlibabaItemPublishMarketGetAPIResponse 从 sync.Pool 获取 AlibabaItemPublishMarketGetAPIResponse
func GetAlibabaItemPublishMarketGetAPIResponse() *AlibabaItemPublishMarketGetAPIResponse {
	return poolAlibabaItemPublishMarketGetAPIResponse.Get().(*AlibabaItemPublishMarketGetAPIResponse)
}

// ReleaseAlibabaItemPublishMarketGetAPIResponse 将 AlibabaItemPublishMarketGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemPublishMarketGetAPIResponse(v *AlibabaItemPublishMarketGetAPIResponse) {
	v.Reset()
	poolAlibabaItemPublishMarketGetAPIResponse.Put(v)
}
