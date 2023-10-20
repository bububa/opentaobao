package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangerefuse 卖家拒绝换货申请
// tmall.exchange.refuse
//
// 卖家拒绝换货申请
func Tmallexchangerefuse(clt *core.SDKClient, req *exchange.TmallexchangerefuseAPIRequest, session string) (*exchange.TmallexchangerefuseAPIResponse, error) {
	var resp exchange.TmallexchangerefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
