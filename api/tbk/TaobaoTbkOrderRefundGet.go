package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkorderrefundget 淘宝客-推广者-全量维权退款订单查询
// taobao.tbk.order.refund.get
//
// 淘宝客账户下全量维权退款订单查询
func Taobaotbkorderrefundget(clt *core.SDKClient, req *tbk.TaobaotbkorderrefundgetAPIRequest, session string) (*tbk.TaobaotbkorderrefundgetAPIResponse, error) {
	var resp tbk.TaobaotbkorderrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
