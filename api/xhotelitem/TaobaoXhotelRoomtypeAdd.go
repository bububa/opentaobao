package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

/* TaobaoXhotelRoomtypeAdd
房型新增接口（ID重复变更新）
taobao.xhotel.roomtype.add

房型添加或更新 */
func TaobaoXhotelRoomtypeAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelRoomtypeAddAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRoomtypeAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
