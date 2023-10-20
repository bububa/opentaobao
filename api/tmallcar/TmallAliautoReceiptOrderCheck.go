package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoreceiptordercheck 查看工单查询订单是否已付款
// tmall.aliauto.receipt.order.check
//
// 查看工单查询订单是否已付款
func Tmallaliautoreceiptordercheck(clt *core.SDKClient, req *tmallcar.TmallaliautoreceiptordercheckAPIRequest, session string) (*tmallcar.TmallaliautoreceiptordercheckAPIResponse, error) {
	var resp tmallcar.TmallaliautoreceiptordercheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
