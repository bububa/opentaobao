package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderAcceptAPIResponse 确认接单 API返回值
// taobao.alitrip.car.order.accept
//
// 用来接收服务商确认接单信息
type TaobaoAlitripCarOrderAcceptAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripCarOrderAcceptAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderAcceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripCarOrderAcceptAPIResponseModel).Reset()
}

// TaobaoAlitripCarOrderAcceptAPIResponseModel is 确认接单 成功返回结果
type TaobaoAlitripCarOrderAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_car_order_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *TaobaoAlitripCarOrderAcceptApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripCarOrderAcceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripCarOrderAcceptAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripCarOrderAcceptAPIResponse)
	},
}

// GetTaobaoAlitripCarOrderAcceptAPIResponse 从 sync.Pool 获取 TaobaoAlitripCarOrderAcceptAPIResponse
func GetTaobaoAlitripCarOrderAcceptAPIResponse() *TaobaoAlitripCarOrderAcceptAPIResponse {
	return poolTaobaoAlitripCarOrderAcceptAPIResponse.Get().(*TaobaoAlitripCarOrderAcceptAPIResponse)
}

// ReleaseTaobaoAlitripCarOrderAcceptAPIResponse 将 TaobaoAlitripCarOrderAcceptAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripCarOrderAcceptAPIResponse(v *TaobaoAlitripCarOrderAcceptAPIResponse) {
	v.Reset()
	poolTaobaoAlitripCarOrderAcceptAPIResponse.Put(v)
}
