package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorprefundreview 审核退款单
// taobao.rp.refund.review
//
// 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
func Taobaorprefundreview(clt *core.SDKClient, req *tbrefund.TaobaorprefundreviewAPIRequest, session string) (*tbrefund.TaobaorprefundreviewAPIResponse, error) {
	var resp tbrefund.TaobaorprefundreviewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
