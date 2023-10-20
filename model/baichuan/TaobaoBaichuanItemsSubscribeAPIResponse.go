package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemsSubscribeAPIResponse 百川批量商品订阅 API返回值
// taobao.baichuan.items.subscribe
//
// 百川批量添加订阅的商品
type TaobaoBaichuanItemsSubscribeAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanItemsSubscribeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemsSubscribeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanItemsSubscribeAPIResponseModel).Reset()
}

// TaobaoBaichuanItemsSubscribeAPIResponseModel is 百川批量商品订阅 成功返回结果
type TaobaoBaichuanItemsSubscribeAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_items_subscribe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoBaichuanItemsSubscribeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemsSubscribeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaichuanItemsSubscribeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemsSubscribeAPIResponse)
	},
}

// GetTaobaoBaichuanItemsSubscribeAPIResponse 从 sync.Pool 获取 TaobaoBaichuanItemsSubscribeAPIResponse
func GetTaobaoBaichuanItemsSubscribeAPIResponse() *TaobaoBaichuanItemsSubscribeAPIResponse {
	return poolTaobaoBaichuanItemsSubscribeAPIResponse.Get().(*TaobaoBaichuanItemsSubscribeAPIResponse)
}

// ReleaseTaobaoBaichuanItemsSubscribeAPIResponse 将 TaobaoBaichuanItemsSubscribeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanItemsSubscribeAPIResponse(v *TaobaoBaichuanItemsSubscribeAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanItemsSubscribeAPIResponse.Put(v)
}
