package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundConfirmreturnAPIResponse 【机票代理商】确认退票 API返回值
// taobao.alitrip.seller.refund.confirmreturn
//
// 确认退票
type TaobaoAlitripSellerRefundConfirmreturnAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundConfirmreturnAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundConfirmreturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerRefundConfirmreturnAPIResponseModel).Reset()
}

// TaobaoAlitripSellerRefundConfirmreturnAPIResponseModel is 【机票代理商】确认退票 成功返回结果
type TaobaoAlitripSellerRefundConfirmreturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refund_confirmreturn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundConfirmreturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoAlitripSellerRefundConfirmreturnAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundConfirmreturnAPIResponse)
	},
}

// GetTaobaoAlitripSellerRefundConfirmreturnAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerRefundConfirmreturnAPIResponse
func GetTaobaoAlitripSellerRefundConfirmreturnAPIResponse() *TaobaoAlitripSellerRefundConfirmreturnAPIResponse {
	return poolTaobaoAlitripSellerRefundConfirmreturnAPIResponse.Get().(*TaobaoAlitripSellerRefundConfirmreturnAPIResponse)
}

// ReleaseTaobaoAlitripSellerRefundConfirmreturnAPIResponse 将 TaobaoAlitripSellerRefundConfirmreturnAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerRefundConfirmreturnAPIResponse(v *TaobaoAlitripSellerRefundConfirmreturnAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerRefundConfirmreturnAPIResponse.Put(v)
}
