package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest
负卖库存失效接口 API请求
alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate

失效负卖库存数据 */
type AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest struct {
	model.Params
	// 入参
	_futureInventoryMainOperationQuest *Futureinventorymainoperationquest
}

// New
