package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelrateget 酒店产品库rate查询
// taobao.xhotel.rate.get
//
// 酒店产品库rate查询
func Taobaoxhotelrateget(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelrategetAPIRequest, session string) (*xhotelitem.TaobaoxhotelrategetAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelrategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
