package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelxitemdelete 删除 x 元素
// taobao.xhotel.xitem.delete
//
// 删除 x 元素
func Taobaoxhotelxitemdelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelxitemdeleteAPIRequest, session string) (*xhotelitem.TaobaoxhotelxitemdeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelxitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
