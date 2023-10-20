package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoBind 自促活动绑定接口
// taobao.xhotel.bnbpromo.bind
//
// 自促活动绑定接口
func TaobaoXhotelBnbpromoBind(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoBindAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbpromoBindAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbpromoBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
