package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbaseinforoomget 酒店房型与房价查询
// taobao.xhotel.baseinfo.room.get
//
// 根据outHid/hid获取酒店的房型和价格信息
func Taobaoxhotelbaseinforoomget(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbaseinforoomgetAPIRequest, session string) (*xhotelitem.TaobaoxhotelbaseinforoomgetAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbaseinforoomgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
