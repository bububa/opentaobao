package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelrateupdate 价格推送接口（全量更新）
// taobao.xhotel.rate.update
//
// 酒店产品库rate更新
func Taobaoxhotelrateupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelrateupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelrateupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelrateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
