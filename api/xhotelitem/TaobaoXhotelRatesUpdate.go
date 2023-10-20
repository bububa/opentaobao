package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelratesupdate 价格推送接口（批量全量）
// taobao.xhotel.rates.update
//
// 酒店产品库rate批量更新房态信息
func Taobaoxhotelratesupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelratesupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelratesupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelratesupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
