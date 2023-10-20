package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorprefundsagree 同意退款
// taobao.rp.refunds.agree
//
// 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
func Taobaorprefundsagree(clt *core.SDKClient, req *tbrefund.TaobaorprefundsagreeAPIRequest, session string) (*tbrefund.TaobaorprefundsagreeAPIResponse, error) {
	var resp tbrefund.TaobaorprefundsagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
