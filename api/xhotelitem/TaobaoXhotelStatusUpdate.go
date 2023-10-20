package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelstatusupdate 酒店状态更新
// taobao.xhotel.status.update
//
// 酒店状态更新
func Taobaoxhotelstatusupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelstatusupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelstatusupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
