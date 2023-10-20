package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefundmessagesget 查询退款留言/凭证列表
// taobao.refund.messages.get
//
// 查询退款留言/凭证列表
func Taobaorefundmessagesget(clt *core.SDKClient, req *tbrefund.TaobaorefundmessagesgetAPIRequest, session string) (*tbrefund.TaobaorefundmessagesgetAPIResponse, error) {
	var resp tbrefund.TaobaorefundmessagesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
