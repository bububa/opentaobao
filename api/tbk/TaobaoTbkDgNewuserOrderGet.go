package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

/* TaobaoTbkDgNewuserOrderGet
淘宝客-推广者-新用户订单明细查询
taobao.tbk.dg.newuser.order.get

拉新API */
func TaobaoTbkDgNewuserOrderGet(clt *core.SDKClient, req *tbk.TaobaoTbkDgNewuserOrderGetAPIRequest, session string) (*tbk.TaobaoTbkDgNewuserOrderGetAPIResponse, error) {
	var resp tbk.TaobaoTbkDgNewuserOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
