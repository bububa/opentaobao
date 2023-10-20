package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkorderdetailsget 淘宝客-推广者-所有订单查询
// taobao.tbk.order.details.get
//
// 淘宝客推广带来的所有拍下付款的正向订单明细报表。
func Taobaotbkorderdetailsget(clt *core.SDKClient, req *tbk.TaobaotbkorderdetailsgetAPIRequest, session string) (*tbk.TaobaotbkorderdetailsgetAPIResponse, error) {
	var resp tbk.TaobaotbkorderdetailsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
