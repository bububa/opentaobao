package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝退款 APIResponse
taobao.refund.refuse

卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：<br/>1. 传入的refund_id和相应的tid, oid必须匹配<br/>2. 如果一笔订单只有一笔子订单，则tid必须与oid相同<br/>3. 只有卖家才能执行拒绝退款操作<br/>4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
*/
type TaobaoRefundRefuseAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRefundRefuseResponse `json:"refund_refuse_response,omitempty"` 
    TaobaoRefundRefuseResponse
}

/* model for simplify = false
type TaobaoRefundRefuseResponse struct {

    // 拒绝退款成功后，会返回Refund数据结构中的refund_id, status, modified字段
    
    Refund  *struct {
        Refund  *Refund `json:"refund,omitempty"`
    } `json:"refund,omitempty"`
    

    // 拒绝退款操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoRefundRefuseResponse struct {

    // 拒绝退款成功后，会返回Refund数据结构中的refund_id, status, modified字段
    Refund   *Refund `json:"refund,omitempty"`

    // 拒绝退款操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
