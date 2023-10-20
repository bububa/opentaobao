package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemsUnsubscribeAPIResponse 批量删除商品订阅 API返回值
// taobao.baichuan.items.unsubscribe
//
// 批量删除商品订阅
type TaobaoBaichuanItemsUnsubscribeAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanItemsUnsubscribeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemsUnsubscribeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanItemsUnsubscribeAPIResponseModel).Reset()
}

// TaobaoBaichuanItemsUnsubscribeAPIResponseModel is 批量删除商品订阅 成功返回结果
type TaobaoBaichuanItemsUnsubscribeAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_items_unsubscribe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoBaichuanItemsUnsubscribeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemsUnsubscribeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaichuanItemsUnsubscribeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemsUnsubscribeAPIResponse)
	},
}

// GetTaobaoBaichuanItemsUnsubscribeAPIResponse 从 sync.Pool 获取 TaobaoBaichuanItemsUnsubscribeAPIResponse
func GetTaobaoBaichuanItemsUnsubscribeAPIResponse() *TaobaoBaichuanItemsUnsubscribeAPIResponse {
	return poolTaobaoBaichuanItemsUnsubscribeAPIResponse.Get().(*TaobaoBaichuanItemsUnsubscribeAPIResponse)
}

// ReleaseTaobaoBaichuanItemsUnsubscribeAPIResponse 将 TaobaoBaichuanItemsUnsubscribeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanItemsUnsubscribeAPIResponse(v *TaobaoBaichuanItemsUnsubscribeAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanItemsUnsubscribeAPIResponse.Put(v)
}
