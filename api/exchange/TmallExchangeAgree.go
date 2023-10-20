package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangeagree 卖家同意换货申请
// tmall.exchange.agree
//
// 卖家同意换货申请
func Tmallexchangeagree(clt *core.SDKClient, req *exchange.TmallexchangeagreeAPIRequest, session string) (*exchange.TmallexchangeagreeAPIResponse, error) {
	var resp exchange.TmallexchangeagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
