package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoroomtypestatusupdate top房型状态修改
// taobao.roomtype.status.update
//
// top房型状态修改
func Taobaoroomtypestatusupdate(clt *core.SDKClient, req *xhotelitem.TaobaoroomtypestatusupdateAPIRequest, session string) (*xhotelitem.TaobaoroomtypestatusupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoroomtypestatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
