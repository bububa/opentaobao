package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelroomtypeadd 房型新增接口（ID重复变更新）
// taobao.xhotel.roomtype.add
//
// 房型添加或更新
func Taobaoxhotelroomtypeadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelroomtypeaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelroomtypeaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelroomtypeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
