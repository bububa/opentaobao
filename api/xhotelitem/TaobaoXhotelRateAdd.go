package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelrateadd 新增专享房价
// taobao.xhotel.rate.add
//
// 酒店产品库rate添加
func Taobaoxhotelrateadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelrateaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelrateaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelrateaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
