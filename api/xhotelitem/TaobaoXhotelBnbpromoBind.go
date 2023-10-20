package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbpromobind 自促活动绑定接口
// taobao.xhotel.bnbpromo.bind
//
// 自促活动绑定接口
func Taobaoxhotelbnbpromobind(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbpromobindAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbpromobindAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbpromobindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
