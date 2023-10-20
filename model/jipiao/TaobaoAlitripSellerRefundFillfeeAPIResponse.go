package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundFillfeeAPIResponse 机票代理商】回填手续费 API返回值
// taobao.alitrip.seller.refund.fillfee
//
// 回填手续费
type TaobaoAlitripSellerRefundFillfeeAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundFillfeeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundFillfeeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerRefundFillfeeAPIResponseModel).Reset()
}

// TaobaoAlitripSellerRefundFillfeeAPIResponseModel is 机票代理商】回填手续费 成功返回结果
type TaobaoAlitripSellerRefundFillfeeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refund_fillfee_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundFillfeeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoAlitripSellerRefundFillfeeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundFillfeeAPIResponse)
	},
}

// GetTaobaoAlitripSellerRefundFillfeeAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerRefundFillfeeAPIResponse
func GetTaobaoAlitripSellerRefundFillfeeAPIResponse() *TaobaoAlitripSellerRefundFillfeeAPIResponse {
	return poolTaobaoAlitripSellerRefundFillfeeAPIResponse.Get().(*TaobaoAlitripSellerRefundFillfeeAPIResponse)
}

// ReleaseTaobaoAlitripSellerRefundFillfeeAPIResponse 将 TaobaoAlitripSellerRefundFillfeeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerRefundFillfeeAPIResponse(v *TaobaoAlitripSellerRefundFillfeeAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerRefundFillfeeAPIResponse.Put(v)
}
