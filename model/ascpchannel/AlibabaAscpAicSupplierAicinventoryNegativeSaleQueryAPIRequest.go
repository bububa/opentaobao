package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest
商家查询负卖库存 API请求
alibaba.ascp.aic.supplier.aicinventory.negative.sale.query

商家根据当前接口查询负卖货品的库存 */
type AlibabaAscpAicSupplierAicinventoryNegativeSaleQueryAPIRequest struct {
	model.Params
	// 库存查询参数
	_aicinventoryQueryRequest *Aicinventoryqueryrequest
}

// New
