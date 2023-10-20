package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemDeleteAPIResponse 全渠道商品删除 API返回值
// taobao.omniitem.item.delete
//
// 全渠道商品删除，能够对门店商品库商品进行删除动作
type TaobaoOmniitemItemDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoOmniitemItemDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniitemItemDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniitemItemDeleteAPIResponseModel).Reset()
}

// TaobaoOmniitemItemDeleteAPIResponseModel is 全渠道商品删除 成功返回结果
type TaobaoOmniitemItemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OmniResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniitemItemDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniitemItemDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniitemItemDeleteAPIResponse)
	},
}

// GetTaobaoOmniitemItemDeleteAPIResponse 从 sync.Pool 获取 TaobaoOmniitemItemDeleteAPIResponse
func GetTaobaoOmniitemItemDeleteAPIResponse() *TaobaoOmniitemItemDeleteAPIResponse {
	return poolTaobaoOmniitemItemDeleteAPIResponse.Get().(*TaobaoOmniitemItemDeleteAPIResponse)
}

// ReleaseTaobaoOmniitemItemDeleteAPIResponse 将 TaobaoOmniitemItemDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniitemItemDeleteAPIResponse(v *TaobaoOmniitemItemDeleteAPIResponse) {
	v.Reset()
	poolTaobaoOmniitemItemDeleteAPIResponse.Put(v)
}
