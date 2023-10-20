package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniOrderGoodsReadyAPIResponse 备货完成 API返回值
// taobao.omni.order.goods.ready
//
// 备货完成
type TaobaoOmniOrderGoodsReadyAPIResponse struct {
	model.CommonResponse
	TaobaoOmniOrderGoodsReadyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniOrderGoodsReadyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniOrderGoodsReadyAPIResponseModel).Reset()
}

// TaobaoOmniOrderGoodsReadyAPIResponseModel is 备货完成 成功返回结果
type TaobaoOmniOrderGoodsReadyAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_order_goods_ready_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniOrderGoodsReadyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniOrderGoodsReadyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniOrderGoodsReadyAPIResponse)
	},
}

// GetTaobaoOmniOrderGoodsReadyAPIResponse 从 sync.Pool 获取 TaobaoOmniOrderGoodsReadyAPIResponse
func GetTaobaoOmniOrderGoodsReadyAPIResponse() *TaobaoOmniOrderGoodsReadyAPIResponse {
	return poolTaobaoOmniOrderGoodsReadyAPIResponse.Get().(*TaobaoOmniOrderGoodsReadyAPIResponse)
}

// ReleaseTaobaoOmniOrderGoodsReadyAPIResponse 将 TaobaoOmniOrderGoodsReadyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniOrderGoodsReadyAPIResponse(v *TaobaoOmniOrderGoodsReadyAPIResponse) {
	v.Reset()
	poolTaobaoOmniOrderGoodsReadyAPIResponse.Put(v)
}
