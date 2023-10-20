package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneOrderExternalCreateAPIResponse 数字虚拟话费外放下单接口 API返回值
// taobao.phone.order.external.create
//
// 数字虚拟话费外放下单接口
type TaobaoPhoneOrderExternalCreateAPIResponse struct {
	model.CommonResponse
	TaobaoPhoneOrderExternalCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPhoneOrderExternalCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPhoneOrderExternalCreateAPIResponseModel).Reset()
}

// TaobaoPhoneOrderExternalCreateAPIResponseModel is 数字虚拟话费外放下单接口 成功返回结果
type TaobaoPhoneOrderExternalCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"phone_order_external_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPhoneOrderExternalCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPhoneOrderExternalCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPhoneOrderExternalCreateAPIResponse)
	},
}

// GetTaobaoPhoneOrderExternalCreateAPIResponse 从 sync.Pool 获取 TaobaoPhoneOrderExternalCreateAPIResponse
func GetTaobaoPhoneOrderExternalCreateAPIResponse() *TaobaoPhoneOrderExternalCreateAPIResponse {
	return poolTaobaoPhoneOrderExternalCreateAPIResponse.Get().(*TaobaoPhoneOrderExternalCreateAPIResponse)
}

// ReleaseTaobaoPhoneOrderExternalCreateAPIResponse 将 TaobaoPhoneOrderExternalCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPhoneOrderExternalCreateAPIResponse(v *TaobaoPhoneOrderExternalCreateAPIResponse) {
	v.Reset()
	poolTaobaoPhoneOrderExternalCreateAPIResponse.Put(v)
}
