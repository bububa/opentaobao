package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundorderlistFetchAPIResponse
【机票代理商】——退票订单列表提取 API返回值
taobao.alitrip.seller.refundorderlist.fetch

代理商纬度退票订单列表提取 */
type TaobaoAlitripSellerRefundorderlistFetchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel
}

// TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel is 【机票代理商】——退票订单列表提取 成功返回结果
type TaobaoAlitripSellerRefundorderlistFetchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refundorderlist_fetch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退票订单列表
	ResultList []ReturnApplyDo `json:"result_list,omitempty" xml:"result_list>return_apply_do,omitempty"`
}
