package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelroomtypeupdate 房型更新接口（ID不存在自动新增）
// taobao.xhotel.roomtype.update
//
// 酒店房型更新或添加
func Taobaoxhotelroomtypeupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelroomtypeupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelroomtypeupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelroomtypeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
