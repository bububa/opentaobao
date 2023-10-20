package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniDealerOdersRefundAddressAPIResponse 经销商查询逆向退货地址 API返回值
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
type TaobaoOmniDealerOdersRefundAddressAPIResponse struct {
	model.CommonResponse
	TaobaoOmniDealerOdersRefundAddressAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniDealerOdersRefundAddressAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniDealerOdersRefundAddressAPIResponseModel).Reset()
}

// TaobaoOmniDealerOdersRefundAddressAPIResponseModel is 经销商查询逆向退货地址 成功返回结果
type TaobaoOmniDealerOdersRefundAddressAPIResponseModel struct {
	XMLName xml.Name `xml:"omni_dealer_oders_refund_address_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 经销商订单退货地址
	Result *TaobaoOmniDealerOdersRefundAddressResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniDealerOdersRefundAddressAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOmniDealerOdersRefundAddressAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniDealerOdersRefundAddressAPIResponse)
	},
}

// GetTaobaoOmniDealerOdersRefundAddressAPIResponse 从 sync.Pool 获取 TaobaoOmniDealerOdersRefundAddressAPIResponse
func GetTaobaoOmniDealerOdersRefundAddressAPIResponse() *TaobaoOmniDealerOdersRefundAddressAPIResponse {
	return poolTaobaoOmniDealerOdersRefundAddressAPIResponse.Get().(*TaobaoOmniDealerOdersRefundAddressAPIResponse)
}

// ReleaseTaobaoOmniDealerOdersRefundAddressAPIResponse 将 TaobaoOmniDealerOdersRefundAddressAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniDealerOdersRefundAddressAPIResponse(v *TaobaoOmniDealerOdersRefundAddressAPIResponse) {
	v.Reset()
	poolTaobaoOmniDealerOdersRefundAddressAPIResponse.Put(v)
}
