package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkrelationrefund 淘宝客-推广者-维权退款订单查询
// taobao.tbk.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
func Taobaotbkrelationrefund(clt *core.SDKClient, req *tbk.TaobaotbkrelationrefundAPIRequest, session string) (*tbk.TaobaotbkrelationrefundAPIResponse, error) {
	var resp tbk.TaobaotbkrelationrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
