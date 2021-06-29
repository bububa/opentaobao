package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收退货V2 API返回值 
taobao.idle.recycle.refund.returngoods

回收商买家退货，填写退货运单号
*/
type TaobaoIdleRecycleRefundReturngoodsAPIResponse struct {
    model.CommonResponse
    TaobaoIdleRecycleRefundReturngoodsResponse
}

// 闲鱼回收退货V2 成功返回结果
type TaobaoIdleRecycleRefundReturngoodsResponse struct {
    XMLName xml.Name `xml:"idle_recycle_refund_returngoods_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 退货
    Result   *IdleTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
