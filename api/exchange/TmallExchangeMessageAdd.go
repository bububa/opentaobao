package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangemessageadd 卖家创建换货留言
// tmall.exchange.message.add
//
// 卖家创建换货留言
func Tmallexchangemessageadd(clt *core.SDKClient, req *exchange.TmallexchangemessageaddAPIRequest, session string) (*exchange.TmallexchangemessageaddAPIResponse, error) {
	var resp exchange.TmallexchangemessageaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
