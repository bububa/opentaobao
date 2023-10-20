package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelrateplanupdate 价格计划rateplan更新或添加
// taobao.xhotel.rateplan.update
//
// 酒店产品库rateplan更新或添加
func Taobaoxhotelrateplanupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelrateplanupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelrateplanupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelrateplanupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
