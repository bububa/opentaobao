package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoUnbind 自促活动解绑接口
// taobao.xhotel.bnbpromo.unbind
//
// 自促活动解绑接口
func TaobaoXhotelBnbpromoUnbind(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoUnbindAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoUnbindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
