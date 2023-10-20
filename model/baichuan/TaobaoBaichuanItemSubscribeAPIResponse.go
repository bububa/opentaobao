package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemSubscribeAPIResponse 单个商品订阅 API返回值
// taobao.baichuan.item.subscribe
//
// 百川单个商品订阅
type TaobaoBaichuanItemSubscribeAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanItemSubscribeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemSubscribeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanItemSubscribeAPIResponseModel).Reset()
}

// TaobaoBaichuanItemSubscribeAPIResponseModel is 单个商品订阅 成功返回结果
type TaobaoBaichuanItemSubscribeAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_item_subscribe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoBaichuanItemSubscribeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemSubscribeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaichuanItemSubscribeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemSubscribeAPIResponse)
	},
}

// GetTaobaoBaichuanItemSubscribeAPIResponse 从 sync.Pool 获取 TaobaoBaichuanItemSubscribeAPIResponse
func GetTaobaoBaichuanItemSubscribeAPIResponse() *TaobaoBaichuanItemSubscribeAPIResponse {
	return poolTaobaoBaichuanItemSubscribeAPIResponse.Get().(*TaobaoBaichuanItemSubscribeAPIResponse)
}

// ReleaseTaobaoBaichuanItemSubscribeAPIResponse 将 TaobaoBaichuanItemSubscribeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanItemSubscribeAPIResponse(v *TaobaoBaichuanItemSubscribeAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanItemSubscribeAPIResponse.Put(v)
}
