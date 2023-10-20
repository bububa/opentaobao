package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhoteladd 酒店新增接口（ID重复自动更新）
// taobao.xhotel.add
//
// 添加酒店或更新酒店
func Taobaoxhoteladd(clt *core.SDKClient, req *xhotelitem.TaobaoxhoteladdAPIRequest, session string) (*xhotelitem.TaobaoxhoteladdAPIResponse, error) {
	var resp xhotelitem.TaobaoxhoteladdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
