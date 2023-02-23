package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpRefundReview 审核退款单
// taobao.rp.refund.review
//
// 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
func TaobaoRpRefundReview(clt *core.SDKClient, req *tbrefund.TaobaoRpRefundReviewAPIRequest, session string) (*tbrefund.TaobaoRpRefundReviewAPIResponse, error) {
	var resp tbrefund.TaobaoRpRefundReviewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
