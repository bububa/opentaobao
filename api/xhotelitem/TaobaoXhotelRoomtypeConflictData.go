package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelroomtypeconflictdata 商家床型冲突数据接口
// taobao.xhotel.roomtype.conflict.data
//
// 商家床型冲突数据接口
func Taobaoxhotelroomtypeconflictdata(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelroomtypeconflictdataAPIRequest, session string) (*xhotelitem.TaobaoxhotelroomtypeconflictdataAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelroomtypeconflictdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
