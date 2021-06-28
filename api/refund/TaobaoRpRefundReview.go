package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
审核退款单 
taobao.rp.refund.review

审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
*/
func TaobaoRpRefundReview(clt *core.SDKClient, req *refund.TaobaoRpRefundReviewRequest, session string) (*refund.TaobaoRpRefundReviewAPIResponse, error) {
    var resp refund.TaobaoRpRefundReviewAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
