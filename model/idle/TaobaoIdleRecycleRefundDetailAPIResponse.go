package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收退款详情V2 API返回值 
taobao.idle.recycle.refund.detail

回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息
*/
type TaobaoIdleRecycleRefundDetailAPIResponse struct {
    model.CommonResponse
    TaobaoIdleRecycleRefundDetailAPIResponseModel
}

// 闲鱼回收退款详情V2 成功返回结果
type TaobaoIdleRecycleRefundDetailAPIResponseModel struct {
    XMLName xml.Name `xml:"idle_recycle_refund_detail_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 退款详情，说明文档：https://www.yuque.com/tushuguanyuan/fsgl7y/xn2lu8#ed2c2d6a
    Result   *IdleTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
