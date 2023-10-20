package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelroomupdate 房型库存推送接口（全量推送）
// taobao.xhotel.room.update
//
// 此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。
func Taobaoxhotelroomupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelroomupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelroomupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelroomupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
