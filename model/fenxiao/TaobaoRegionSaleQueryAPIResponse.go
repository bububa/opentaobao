package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionSaleQueryAPIResponse 查询商品销售区域 API返回值
// taobao.region.sale.query
//
// 查询商品销售区域
type TaobaoRegionSaleQueryAPIResponse struct {
	model.CommonResponse
	TaobaoRegionSaleQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRegionSaleQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRegionSaleQueryAPIResponseModel).Reset()
}

// TaobaoRegionSaleQueryAPIResponseModel is 查询商品销售区域 成功返回结果
type TaobaoRegionSaleQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"region_sale_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRegionSaleQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRegionSaleQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRegionSaleQueryAPIResponse)
	},
}

// GetTaobaoRegionSaleQueryAPIResponse 从 sync.Pool 获取 TaobaoRegionSaleQueryAPIResponse
func GetTaobaoRegionSaleQueryAPIResponse() *TaobaoRegionSaleQueryAPIResponse {
	return poolTaobaoRegionSaleQueryAPIResponse.Get().(*TaobaoRegionSaleQueryAPIResponse)
}

// ReleaseTaobaoRegionSaleQueryAPIResponse 将 TaobaoRegionSaleQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRegionSaleQueryAPIResponse(v *TaobaoRegionSaleQueryAPIResponse) {
	v.Reset()
	poolTaobaoRegionSaleQueryAPIResponse.Put(v)
}
