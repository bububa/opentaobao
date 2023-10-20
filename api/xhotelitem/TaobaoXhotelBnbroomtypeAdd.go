package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbroomtypeadd 民宿新增房源
// taobao.xhotel.bnbroomtype.add
//
// 添加民宿房源
func Taobaoxhotelbnbroomtypeadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbroomtypeaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbroomtypeaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbroomtypeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
