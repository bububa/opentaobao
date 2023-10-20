package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundMessagesGet 查询退款留言/凭证列表
// taobao.refund.messages.get
//
// 查询退款留言/凭证列表
func TaobaoRefundMessagesGet(clt *core.SDKClient, req *tbrefund.TaobaoRefundMessagesGetAPIRequest, resp *tbrefund.TaobaoRefundMessagesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
