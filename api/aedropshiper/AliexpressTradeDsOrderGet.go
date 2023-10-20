package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpresstradedsorderget 买家查询订单详情
// aliexpress.trade.ds.order.get
//
// 买家查询订单详情，用于dropshipper
func Aliexpresstradedsorderget(clt *core.SDKClient, req *aedropshiper.AliexpresstradedsordergetAPIRequest, session string) (*aedropshiper.AliexpresstradedsordergetAPIResponse, error) {
	var resp aedropshiper.AliexpresstradedsordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
