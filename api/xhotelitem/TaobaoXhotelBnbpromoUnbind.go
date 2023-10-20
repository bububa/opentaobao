package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbpromounbind 自促活动解绑接口
// taobao.xhotel.bnbpromo.unbind
//
// 自促活动解绑接口
func Taobaoxhotelbnbpromounbind(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbpromounbindAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbpromounbindAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbpromounbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
