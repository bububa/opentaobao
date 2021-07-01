package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest
AIC负卖库存新增和修改接口 API请求
alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish

新增负卖库存记录和变更负卖库存记录 */
type AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest struct {
	model.Params
	// 入参
	_futureInventoryMainOperationQuest *Futureinventorymainoperationquest
}

// New
