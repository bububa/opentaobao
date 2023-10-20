package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgpunishorderget 淘宝客-推广者-处罚订单查询
// taobao.tbk.dg.punish.order.get
//
// 新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的
func Taobaotbkdgpunishorderget(clt *core.SDKClient, req *tbk.TaobaotbkdgpunishordergetAPIRequest, session string) (*tbk.TaobaotbkdgpunishordergetAPIResponse, error) {
	var resp tbk.TaobaotbkdgpunishordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
