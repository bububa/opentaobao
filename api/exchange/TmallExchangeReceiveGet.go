package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangereceiveget 卖家查询换货列表
// tmall.exchange.receive.get
//
// 卖家查询换货列表
func Tmallexchangereceiveget(clt *core.SDKClient, req *exchange.TmallexchangereceivegetAPIRequest, session string) (*exchange.TmallexchangereceivegetAPIResponse, error) {
	var resp exchange.TmallexchangereceivegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
