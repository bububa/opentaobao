package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangemessagesget 查询换货订单留言列表
// tmall.exchange.messages.get
//
// 查询换货订单留言列表
func Tmallexchangemessagesget(clt *core.SDKClient, req *exchange.TmallexchangemessagesgetAPIRequest, session string) (*exchange.TmallexchangemessagesgetAPIResponse, error) {
	var resp exchange.TmallexchangemessagesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
