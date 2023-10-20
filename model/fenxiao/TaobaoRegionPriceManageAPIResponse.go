package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionPriceManageAPIResponse 编辑区域价格 API返回值
// taobao.region.price.manage
//
// 编辑区域价格
type TaobaoRegionPriceManageAPIResponse struct {
	model.CommonResponse
	TaobaoRegionPriceManageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRegionPriceManageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRegionPriceManageAPIResponseModel).Reset()
}

// TaobaoRegionPriceManageAPIResponseModel is 编辑区域价格 成功返回结果
type TaobaoRegionPriceManageAPIResponseModel struct {
	XMLName xml.Name `xml:"region_price_manage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRegionPriceManageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoRegionPriceManageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRegionPriceManageAPIResponse)
	},
}

// GetTaobaoRegionPriceManageAPIResponse 从 sync.Pool 获取 TaobaoRegionPriceManageAPIResponse
func GetTaobaoRegionPriceManageAPIResponse() *TaobaoRegionPriceManageAPIResponse {
	return poolTaobaoRegionPriceManageAPIResponse.Get().(*TaobaoRegionPriceManageAPIResponse)
}

// ReleaseTaobaoRegionPriceManageAPIResponse 将 TaobaoRegionPriceManageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRegionPriceManageAPIResponse(v *TaobaoRegionPriceManageAPIResponse) {
	v.Reset()
	poolTaobaoRegionPriceManageAPIResponse.Put(v)
}
