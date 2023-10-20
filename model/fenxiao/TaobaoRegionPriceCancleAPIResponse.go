package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionPriceCancleAPIResponse 取消区域价格 API返回值
// taobao.region.price.cancle
//
// 取消区域价格
type TaobaoRegionPriceCancleAPIResponse struct {
	model.CommonResponse
	TaobaoRegionPriceCancleAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRegionPriceCancleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRegionPriceCancleAPIResponseModel).Reset()
}

// TaobaoRegionPriceCancleAPIResponseModel is 取消区域价格 成功返回结果
type TaobaoRegionPriceCancleAPIResponseModel struct {
	XMLName xml.Name `xml:"region_price_cancle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRegionPriceCancleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRegionPriceCancleAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRegionPriceCancleAPIResponse)
	},
}

// GetTaobaoRegionPriceCancleAPIResponse 从 sync.Pool 获取 TaobaoRegionPriceCancleAPIResponse
func GetTaobaoRegionPriceCancleAPIResponse() *TaobaoRegionPriceCancleAPIResponse {
	return poolTaobaoRegionPriceCancleAPIResponse.Get().(*TaobaoRegionPriceCancleAPIResponse)
}

// ReleaseTaobaoRegionPriceCancleAPIResponse 将 TaobaoRegionPriceCancleAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRegionPriceCancleAPIResponse(v *TaobaoRegionPriceCancleAPIResponse) {
	v.Reset()
	poolTaobaoRegionPriceCancleAPIResponse.Put(v)
}
