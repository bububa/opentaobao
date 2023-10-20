package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelupdate 酒店更新接口（ID不存在自动新增）
// taobao.xhotel.update
//
// 酒店更新接口
func Taobaoxhotelupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
