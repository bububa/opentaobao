package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemDeleteAPIResponse 删除单条商品 API返回值
// taobao.item.delete
//
// 删除单条商品
type TaobaoItemDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoItemDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemDeleteAPIResponseModel).Reset()
}

// TaobaoItemDeleteAPIResponseModel is 删除单条商品 成功返回结果
type TaobaoItemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被删除商品的相关信息
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Item = nil
}

var poolTaobaoItemDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemDeleteAPIResponse)
	},
}

// GetTaobaoItemDeleteAPIResponse 从 sync.Pool 获取 TaobaoItemDeleteAPIResponse
func GetTaobaoItemDeleteAPIResponse() *TaobaoItemDeleteAPIResponse {
	return poolTaobaoItemDeleteAPIResponse.Get().(*TaobaoItemDeleteAPIResponse)
}

// ReleaseTaobaoItemDeleteAPIResponse 将 TaobaoItemDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemDeleteAPIResponse(v *TaobaoItemDeleteAPIResponse) {
	v.Reset()
	poolTaobaoItemDeleteAPIResponse.Put(v)
}
