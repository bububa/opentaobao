package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkOrderDetailsGet 淘宝客-推广者-所有订单查询
// taobao.tbk.order.details.get
//
// 淘宝客订单查询
func TaobaoTbkOrderDetailsGet(clt *core.SDKClient, req *tbk.TaobaoTbkOrderDetailsGetAPIRequest, session string) (*tbk.TaobaoTbkOrderDetailsGetAPIResponse, error) {
	var resp tbk.TaobaoTbkOrderDetailsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
