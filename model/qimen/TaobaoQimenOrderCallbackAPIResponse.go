package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderCallbackAPIResponse 配送拦截接口 API返回值
// taobao.qimen.order.callback
//
// 配送拦截
type TaobaoQimenOrderCallbackAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenOrderCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenOrderCallbackAPIResponseModel).Reset()
}

// TaobaoQimenOrderCallbackAPIResponseModel is 配送拦截接口 成功返回结果
type TaobaoQimenOrderCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_order_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *OrderCallbackResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenOrderCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenOrderCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderCallbackAPIResponse)
	},
}

// GetTaobaoQimenOrderCallbackAPIResponse 从 sync.Pool 获取 TaobaoQimenOrderCallbackAPIResponse
func GetTaobaoQimenOrderCallbackAPIResponse() *TaobaoQimenOrderCallbackAPIResponse {
	return poolTaobaoQimenOrderCallbackAPIResponse.Get().(*TaobaoQimenOrderCallbackAPIResponse)
}

// ReleaseTaobaoQimenOrderCallbackAPIResponse 将 TaobaoQimenOrderCallbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenOrderCallbackAPIResponse(v *TaobaoQimenOrderCallbackAPIResponse) {
	v.Reset()
	poolTaobaoQimenOrderCallbackAPIResponse.Put(v)
}
