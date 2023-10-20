package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbinventorylogquery 根据商品ID查询所有库存变更记录
// taobao.wlb.inventorylog.query
//
// 通过商品ID等几个条件来分页查询库存变更记录
func Taobaowlbinventorylogquery(clt *core.SDKClient, req *wlb.TaobaowlbinventorylogqueryAPIRequest, session string) (*wlb.TaobaowlbinventorylogqueryAPIResponse, error) {
	var resp wlb.TaobaowlbinventorylogqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
