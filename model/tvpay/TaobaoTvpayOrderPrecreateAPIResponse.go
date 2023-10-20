package tvpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayOrderPrecreateAPIResponse tv支付预下单 API返回值
// taobao.tvpay.order.precreate
//
// tv支付预下单
type TaobaoTvpayOrderPrecreateAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayOrderPrecreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTvpayOrderPrecreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTvpayOrderPrecreateAPIResponseModel).Reset()
}

// TaobaoTvpayOrderPrecreateAPIResponseModel is tv支付预下单 成功返回结果
type TaobaoTvpayOrderPrecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_order_precreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTvpayOrderPrecreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTvpayOrderPrecreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTvpayOrderPrecreateAPIResponse)
	},
}

// GetTaobaoTvpayOrderPrecreateAPIResponse 从 sync.Pool 获取 TaobaoTvpayOrderPrecreateAPIResponse
func GetTaobaoTvpayOrderPrecreateAPIResponse() *TaobaoTvpayOrderPrecreateAPIResponse {
	return poolTaobaoTvpayOrderPrecreateAPIResponse.Get().(*TaobaoTvpayOrderPrecreateAPIResponse)
}

// ReleaseTaobaoTvpayOrderPrecreateAPIResponse 将 TaobaoTvpayOrderPrecreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTvpayOrderPrecreateAPIResponse(v *TaobaoTvpayOrderPrecreateAPIResponse) {
	v.Reset()
	poolTaobaoTvpayOrderPrecreateAPIResponse.Put(v)
}
