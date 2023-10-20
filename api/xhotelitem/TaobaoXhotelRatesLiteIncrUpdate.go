package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelratesliteincrupdate 酒店价格库存轻量级增量接口
// taobao.xhotel.rates.lite.incr.update
//
// 多个rate的库存房价开关的增量更新接口
func Taobaoxhotelratesliteincrupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelratesliteincrupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelratesliteincrupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelratesliteincrupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
