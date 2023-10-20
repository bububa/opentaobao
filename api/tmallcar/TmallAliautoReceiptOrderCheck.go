package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoReceiptOrderCheck 查看工单查询订单是否已付款
// tmall.aliauto.receipt.order.check
//
// 查看工单查询订单是否已付款
func TmallAliautoReceiptOrderCheck(clt *core.SDKClient, req *tmallcar.TmallAliautoReceiptOrderCheckAPIRequest, session string) (*tmallcar.TmallAliautoReceiptOrderCheckAPIResponse, error) {
	var resp tmallcar.TmallAliautoReceiptOrderCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
