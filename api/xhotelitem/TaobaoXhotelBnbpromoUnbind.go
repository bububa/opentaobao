package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoUnbind 自促活动解绑接口
// taobao.xhotel.bnbpromo.unbind
//
// 自促活动解绑接口
func TaobaoXhotelBnbpromoUnbind(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoUnbindAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbpromoUnbindAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbpromoUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
