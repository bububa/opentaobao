package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundSearchAPIResponse 【机票代理商】退票申请单列表 API返回值
// taobao.alitrip.seller.refund.search
//
// 查询退票申请单列表
type TaobaoAlitripSellerRefundSearchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerRefundSearchAPIResponseModel).Reset()
}

// TaobaoAlitripSellerRefundSearchAPIResponseModel is 【机票代理商】退票申请单列表 成功返回结果
type TaobaoAlitripSellerRefundSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refund_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoAlitripSellerRefundSearchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripSellerRefundSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundSearchAPIResponse)
	},
}

// GetTaobaoAlitripSellerRefundSearchAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerRefundSearchAPIResponse
func GetTaobaoAlitripSellerRefundSearchAPIResponse() *TaobaoAlitripSellerRefundSearchAPIResponse {
	return poolTaobaoAlitripSellerRefundSearchAPIResponse.Get().(*TaobaoAlitripSellerRefundSearchAPIResponse)
}

// ReleaseTaobaoAlitripSellerRefundSearchAPIResponse 将 TaobaoAlitripSellerRefundSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerRefundSearchAPIResponse(v *TaobaoAlitripSellerRefundSearchAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerRefundSearchAPIResponse.Put(v)
}
