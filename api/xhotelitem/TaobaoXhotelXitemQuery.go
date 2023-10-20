package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelxitemquery 查询 x 元素
// taobao.xhotel.xitem.query
//
// 查询 x 元素
func Taobaoxhotelxitemquery(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelxitemqueryAPIRequest, session string) (*xhotelitem.TaobaoxhotelxitemqueryAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelxitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
