package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbreviewadd 对外开放评论接口
// taobao.xhotel.bnbreview.add
//
// 对外开放评论接口
func Taobaoxhotelbnbreviewadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbreviewaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbreviewaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbreviewaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
