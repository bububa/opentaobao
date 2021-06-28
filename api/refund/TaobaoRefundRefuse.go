package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
卖家拒绝退款 
taobao.refund.refuse

卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：<br/>1. 传入的refund_id和相应的tid, oid必须匹配<br/>2. 如果一笔订单只有一笔子订单，则tid必须与oid相同<br/>3. 只有卖家才能执行拒绝退款操作<br/>4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
*/
func TaobaoRefundRefuse(clt *core.SDKClient, req *refund.TaobaoRefundRefuseRequest, session string) (*refund.TaobaoRefundRefuseAPIResponse, error) {
    var resp refund.TaobaoRefundRefuseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
