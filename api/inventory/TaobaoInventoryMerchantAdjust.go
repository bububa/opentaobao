package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// TaobaoInventoryMerchantAdjust 货品库存商家端调整
// taobao.inventory.merchant.adjust
//
// 货品库存商家端调整 ，入库，出库，盘点
func TaobaoInventoryMerchantAdjust(clt *core.SDKClient, req *inventory.TaobaoInventoryMerchantAdjustAPIRequest, session string) (*inventory.TaobaoInventoryMerchantAdjustAPIResponse, error) {
	var resp inventory.TaobaoInventoryMerchantAdjustAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
