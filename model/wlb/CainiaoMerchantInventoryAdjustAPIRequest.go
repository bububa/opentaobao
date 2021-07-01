package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoMerchantInventoryAdjustAPIRequest
商家库存调整 API请求
cainiao.merchant.inventory.adjust

商家仓库存调整接口，目前仅支持全量更新 */
type CainiaoMerchantInventoryAdjustAPIRequest struct {
	model.Params
	// 商家仓编辑库存
	_adjustRequest []MerStoreInvAdjustDto
	// 调用方应用名
	_appName string
	// 操作
	_operation string
}

// New
