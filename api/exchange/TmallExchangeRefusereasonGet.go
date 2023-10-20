package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangerefusereasonget 获取拒绝换货原因列表
// tmall.exchange.refusereason.get
//
// 获取拒绝换货原因列表
func Tmallexchangerefusereasonget(clt *core.SDKClient, req *exchange.TmallexchangerefusereasongetAPIRequest, session string) (*exchange.TmallexchangerefusereasongetAPIResponse, error) {
	var resp exchange.TmallexchangerefusereasongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
