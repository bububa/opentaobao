package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelhouseroomtypeadd 民宿房型信息添加
// taobao.xhotel.house.roomtype.add
//
// 房型添加或更新
func Taobaoxhotelhouseroomtypeadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelhouseroomtypeaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelhouseroomtypeaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelhouseroomtypeaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
