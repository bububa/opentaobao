package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundMessagesGet 查询退款留言/凭证列表
// taobao.refund.messages.get
//
// 查询退款留言/凭证列表
func TaobaoRefundMessagesGet(clt *core.SDKClient, req *tbrefund.TaobaoRefundMessagesGetAPIRequest, session string) (*tbrefund.TaobaoRefundMessagesGetAPIResponse, error) {
	var resp tbrefund.TaobaoRefundMessagesGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
