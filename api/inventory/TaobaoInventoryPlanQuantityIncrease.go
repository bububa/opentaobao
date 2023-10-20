package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// Taobaoinventoryplanquantityincrease 计划库存的增量编辑
// taobao.inventory.plan.quantity.increase
//
// 计划库存的增量编辑
func Taobaoinventoryplanquantityincrease(clt *core.SDKClient, req *inventory.TaobaoinventoryplanquantityincreaseAPIRequest, session string) (*inventory.TaobaoinventoryplanquantityincreaseAPIResponse, error) {
	var resp inventory.TaobaoinventoryplanquantityincreaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
