package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemGetAPIResponse 获取全渠道门店商品 API返回值
// taobao.omniitem.item.get
//
// 通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
type TaobaoOmniitemItemGetAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemItemGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniitemItemGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniitemItemGetAPIResponseModel).Reset()
}

// TaobaoOmniitemItemGetAPIResponseModel is 获取全渠道门店商品 成功返回结果
type TaobaoOmniitemItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniitemItemGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniitemItemGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniitemItemGetAPIResponse)
	},
}

// GetTaobaoOmniitemItemGetAPIResponse 从 sync.Pool 获取 TaobaoOmniitemItemGetAPIResponse
func GetTaobaoOmniitemItemGetAPIResponse() *TaobaoOmniitemItemGetAPIResponse {
	return poolTaobaoOmniitemItemGetAPIResponse.Get().(*TaobaoOmniitemItemGetAPIResponse)
}

// ReleaseTaobaoOmniitemItemGetAPIResponse 将 TaobaoOmniitemItemGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniitemItemGetAPIResponse(v *TaobaoOmniitemItemGetAPIResponse) {
	v.Reset()
	poolTaobaoOmniitemItemGetAPIResponse.Put(v)
}
