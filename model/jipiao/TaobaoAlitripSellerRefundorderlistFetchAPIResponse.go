package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundorderlistFetchAPIResponse 【机票代理商】——退票订单列表提取 API返回值
// taobao.alitrip.seller.refundorderlist.fetch
//
// 代理商纬度退票订单列表提取
type TaobaoAlitripSellerRefundorderlistFetchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundorderlistFetchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel).Reset()
}

// TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel is 【机票代理商】——退票订单列表提取 成功返回结果
type TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refundorderlist_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退票订单列表
	ResultList []ReturnApplyDo `json:"result_list,omitempty" xml:"result_list>return_apply_do,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolTaobaoAlitripSellerRefundorderlistFetchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundorderlistFetchAPIResponse)
	},
}

// GetTaobaoAlitripSellerRefundorderlistFetchAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerRefundorderlistFetchAPIResponse
func GetTaobaoAlitripSellerRefundorderlistFetchAPIResponse() *TaobaoAlitripSellerRefundorderlistFetchAPIResponse {
	return poolTaobaoAlitripSellerRefundorderlistFetchAPIResponse.Get().(*TaobaoAlitripSellerRefundorderlistFetchAPIResponse)
}

// ReleaseTaobaoAlitripSellerRefundorderlistFetchAPIResponse 将 TaobaoAlitripSellerRefundorderlistFetchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerRefundorderlistFetchAPIResponse(v *TaobaoAlitripSellerRefundorderlistFetchAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerRefundorderlistFetchAPIResponse.Put(v)
}
