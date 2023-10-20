package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundGetAPIResponse 【机票代理商】退票申请单详情 API返回值
// taobao.alitrip.seller.refund.get
//
// 查询退票申请单详情
type TaobaoAlitripSellerRefundGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerRefundGetAPIResponseModel).Reset()
}

// TaobaoAlitripSellerRefundGetAPIResponseModel is 【机票代理商】退票申请单详情 成功返回结果
type TaobaoAlitripSellerRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoAlitripSellerRefundGetResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripSellerRefundGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundGetAPIResponse)
	},
}

// GetTaobaoAlitripSellerRefundGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerRefundGetAPIResponse
func GetTaobaoAlitripSellerRefundGetAPIResponse() *TaobaoAlitripSellerRefundGetAPIResponse {
	return poolTaobaoAlitripSellerRefundGetAPIResponse.Get().(*TaobaoAlitripSellerRefundGetAPIResponse)
}

// ReleaseTaobaoAlitripSellerRefundGetAPIResponse 将 TaobaoAlitripSellerRefundGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerRefundGetAPIResponse(v *TaobaoAlitripSellerRefundGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerRefundGetAPIResponse.Put(v)
}
