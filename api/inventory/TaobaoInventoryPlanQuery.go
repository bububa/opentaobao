package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// Taobaoinventoryplanquery 计划库存查询
// taobao.inventory.plan.query
//
// 计划库存查询
func Taobaoinventoryplanquery(clt *core.SDKClient, req *inventory.TaobaoinventoryplanqueryAPIRequest, session string) (*inventory.TaobaoinventoryplanqueryAPIResponse, error) {
	var resp inventory.TaobaoinventoryplanqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
