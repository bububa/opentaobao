package tvpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayOrderQueryAPIResponse tv支付查询订单状态 API返回值
// taobao.tvpay.order.query
//
// tv支付查询订单状态
type TaobaoTvpayOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTvpayOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTvpayOrderQueryAPIResponseModel).Reset()
}

// TaobaoTvpayOrderQueryAPIResponseModel is tv支付查询订单状态 成功返回结果
type TaobaoTvpayOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTvpayOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTvpayOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTvpayOrderQueryAPIResponse)
	},
}

// GetTaobaoTvpayOrderQueryAPIResponse 从 sync.Pool 获取 TaobaoTvpayOrderQueryAPIResponse
func GetTaobaoTvpayOrderQueryAPIResponse() *TaobaoTvpayOrderQueryAPIResponse {
	return poolTaobaoTvpayOrderQueryAPIResponse.Get().(*TaobaoTvpayOrderQueryAPIResponse)
}

// ReleaseTaobaoTvpayOrderQueryAPIResponse 将 TaobaoTvpayOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTvpayOrderQueryAPIResponse(v *TaobaoTvpayOrderQueryAPIResponse) {
	v.Reset()
	poolTaobaoTvpayOrderQueryAPIResponse.Put(v)
}
