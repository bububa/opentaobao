package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelroomtypedeletepublic 商家删除房型数据接口
// taobao.xhotel.roomtype.delete.public
//
// 房型删除TOP接口
func Taobaoxhotelroomtypedeletepublic(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelroomtypedeletepublicAPIRequest, session string) (*xhotelitem.TaobaoxhotelroomtypedeletepublicAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelroomtypedeletepublicAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
