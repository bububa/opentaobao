package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangeget 获取单笔换货详情
// tmall.exchange.get
//
// 获取单笔换货详情
func Tmallexchangeget(clt *core.SDKClient, req *exchange.TmallexchangegetAPIRequest, session string) (*exchange.TmallexchangegetAPIResponse, error) {
	var resp exchange.TmallexchangegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
