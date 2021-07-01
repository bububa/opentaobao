package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoBimTradeorderConsignAPIRequest
驱动保税交易订单发货 API请求
cainiao.bim.tradeorder.consign

驱动保税交易订单发货 */
type CainiaoBimTradeorderConsignAPIRequest struct {
	model.Params
	// 交易单号
	_tradeId string
	// 仓储编码，ERP指定仓库发货时需要传值，编码由菜鸟提供
	_storeCode string
	// 选择的线路ID非必填字段
	_resId string
}

// New
