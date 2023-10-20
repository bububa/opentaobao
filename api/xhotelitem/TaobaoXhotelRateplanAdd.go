package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelrateplanadd 酒店产品库rateplan添加
// taobao.xhotel.rateplan.add
//
// 酒店产品库rateplan
func Taobaoxhotelrateplanadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelrateplanaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelrateplanaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelrateplanaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
