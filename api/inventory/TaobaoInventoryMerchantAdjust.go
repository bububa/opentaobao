package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// Taobaoinventorymerchantadjust 货品库存商家端调整
// taobao.inventory.merchant.adjust
//
// 货品库存商家端调整 ，入库，出库，盘点
func Taobaoinventorymerchantadjust(clt *core.SDKClient, req *inventory.TaobaoinventorymerchantadjustAPIRequest, session string) (*inventory.TaobaoinventorymerchantadjustAPIResponse, error) {
	var resp inventory.TaobaoinventorymerchantadjustAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
