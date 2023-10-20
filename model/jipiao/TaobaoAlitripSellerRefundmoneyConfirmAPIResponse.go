package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundmoneyConfirmAPIResponse 【机票代理商订单】确认退款 API返回值
// taobao.alitrip.seller.refundmoney.confirm
//
// 代理人确认退票申请单的退款
type TaobaoAlitripSellerRefundmoneyConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundmoneyConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundmoneyConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerRefundmoneyConfirmAPIResponseModel).Reset()
}

// TaobaoAlitripSellerRefundmoneyConfirmAPIResponseModel is 【机票代理商订单】确认退款 成功返回结果
type TaobaoAlitripSellerRefundmoneyConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refundmoney_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功确认
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundmoneyConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoAlitripSellerRefundmoneyConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundmoneyConfirmAPIResponse)
	},
}

// GetTaobaoAlitripSellerRefundmoneyConfirmAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerRefundmoneyConfirmAPIResponse
func GetTaobaoAlitripSellerRefundmoneyConfirmAPIResponse() *TaobaoAlitripSellerRefundmoneyConfirmAPIResponse {
	return poolTaobaoAlitripSellerRefundmoneyConfirmAPIResponse.Get().(*TaobaoAlitripSellerRefundmoneyConfirmAPIResponse)
}

// ReleaseTaobaoAlitripSellerRefundmoneyConfirmAPIResponse 将 TaobaoAlitripSellerRefundmoneyConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerRefundmoneyConfirmAPIResponse(v *TaobaoAlitripSellerRefundmoneyConfirmAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerRefundmoneyConfirmAPIResponse.Put(v)
}
