package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelmultipleratesupdate 复杂价格推送接口（批量全量）
// taobao.xhotel.multiplerates.update
//
// 批量更新复杂价格
// 涵盖了taobao.xhotel.rates.update的功能
func Taobaoxhotelmultipleratesupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelmultipleratesupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelmultipleratesupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelmultipleratesupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
