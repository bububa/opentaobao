package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundGet 获取单笔退款详情
// taobao.refund.get
//
// 获取单笔退款详情
func TaobaoRefundGet(clt *core.SDKClient, req *tbrefund.TaobaoRefundGetAPIRequest, resp *tbrefund.TaobaoRefundGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
