package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpressdstradeorderget 交易订单查询
// aliexpress.ds.trade.order.get
//
// 交易订单查询
func Aliexpressdstradeorderget(clt *core.SDKClient, req *aedropshiper.AliexpressdstradeordergetAPIRequest, session string) (*aedropshiper.AliexpressdstradeordergetAPIResponse, error) {
	var resp aedropshiper.AliexpressdstradeordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
