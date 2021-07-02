package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeGet 房型查询接口
// taobao.xhotel.roomtype.get
//
// 房型查询房型查询接口返回结果增加date_confirm字段
func TaobaoXhotelRoomtypeGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelRoomtypeGetAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRoomtypeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
