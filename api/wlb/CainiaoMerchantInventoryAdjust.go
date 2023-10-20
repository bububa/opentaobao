package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// CainiaoMerchantInventoryAdjust 商家库存调整
// cainiao.merchant.inventory.adjust
//
// 商家仓库存调整接口，目前仅支持全量更新
func CainiaoMerchantInventoryAdjust(clt *core.SDKClient, req *wlb.CainiaoMerchantInventoryAdjustAPIRequest, resp *wlb.CainiaoMerchantInventoryAdjustAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
