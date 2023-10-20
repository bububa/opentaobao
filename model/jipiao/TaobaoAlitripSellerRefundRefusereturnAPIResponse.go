package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundRefusereturnAPIResponse 【机票代理商】拒绝退票 API返回值
// taobao.alitrip.seller.refund.refusereturn
//
// 拒绝退票
type TaobaoAlitripSellerRefundRefusereturnAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundRefusereturnAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundRefusereturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerRefundRefusereturnAPIResponseModel).Reset()
}

// TaobaoAlitripSellerRefundRefusereturnAPIResponseModel is 【机票代理商】拒绝退票 成功返回结果
type TaobaoAlitripSellerRefundRefusereturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refund_refusereturn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoAlitripSellerRefundRefusereturnResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundRefusereturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripSellerRefundRefusereturnAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundRefusereturnAPIResponse)
	},
}

// GetTaobaoAlitripSellerRefundRefusereturnAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerRefundRefusereturnAPIResponse
func GetTaobaoAlitripSellerRefundRefusereturnAPIResponse() *TaobaoAlitripSellerRefundRefusereturnAPIResponse {
	return poolTaobaoAlitripSellerRefundRefusereturnAPIResponse.Get().(*TaobaoAlitripSellerRefundRefusereturnAPIResponse)
}

// ReleaseTaobaoAlitripSellerRefundRefusereturnAPIResponse 将 TaobaoAlitripSellerRefundRefusereturnAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerRefundRefusereturnAPIResponse(v *TaobaoAlitripSellerRefundRefusereturnAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerRefundRefusereturnAPIResponse.Put(v)
}
