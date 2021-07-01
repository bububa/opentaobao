package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkStockCabinetQueryAPIRequest
仓内实时库位库存查询 API请求
alibaba.wdk.stock.cabinet.query

查询仓内实时库位库存信息 */
type AlibabaWdkStockCabinetQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_query *WmsInventoryTopQuery
}

// New
