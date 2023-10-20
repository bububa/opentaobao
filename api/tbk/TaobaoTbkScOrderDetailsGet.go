package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscorderdetailsget 淘宝客-服务商-所有订单查询
// taobao.tbk.sc.order.details.get
//
// 服务商使用。可通过该接口查询推广者账号下对应的推广订单明细。
func Taobaotbkscorderdetailsget(clt *core.SDKClient, req *tbk.TaobaotbkscorderdetailsgetAPIRequest, session string) (*tbk.TaobaotbkscorderdetailsgetAPIResponse, error) {
	var resp tbk.TaobaotbkscorderdetailsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
