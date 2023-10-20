package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionPriceQueryAPIResponse 区域价格查询 API返回值
// taobao.region.price.query
//
// 区域价格查询
type TaobaoRegionPriceQueryAPIResponse struct {
	model.CommonResponse
	TaobaoRegionPriceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRegionPriceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRegionPriceQueryAPIResponseModel).Reset()
}

// TaobaoRegionPriceQueryAPIResponseModel is 区域价格查询 成功返回结果
type TaobaoRegionPriceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"region_price_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRegionPriceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRegionPriceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRegionPriceQueryAPIResponse)
	},
}

// GetTaobaoRegionPriceQueryAPIResponse 从 sync.Pool 获取 TaobaoRegionPriceQueryAPIResponse
func GetTaobaoRegionPriceQueryAPIResponse() *TaobaoRegionPriceQueryAPIResponse {
	return poolTaobaoRegionPriceQueryAPIResponse.Get().(*TaobaoRegionPriceQueryAPIResponse)
}

// ReleaseTaobaoRegionPriceQueryAPIResponse 将 TaobaoRegionPriceQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRegionPriceQueryAPIResponse(v *TaobaoRegionPriceQueryAPIResponse) {
	v.Reset()
	poolTaobaoRegionPriceQueryAPIResponse.Put(v)
}
