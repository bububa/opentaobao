package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkStockRealQueryAPIRequest
仓内实时库存查询 API请求
alibaba.wdk.stock.real.query

查询仓内实时库存信息 */
type AlibabaWdkStockRealQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_query *WmsInventoryTopQuery
}

// New
