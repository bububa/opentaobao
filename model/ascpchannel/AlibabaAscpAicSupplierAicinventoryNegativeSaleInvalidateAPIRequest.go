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

// NewAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest 初始化AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest对象
func NewAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest {
	return &AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FutureInventoryMainOperationQuest Setter
// 入参
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) SetFutureInventoryMainOperationQuest(_futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
	r._futureInventoryMainOperationQuest = _futureInventoryMainOperationQuest
	r.Set("future_inventory_main_operation_quest", _futureInventoryMainOperationQuest)
	return nil
}

// Get FutureInventoryMainOperationQuest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateAPIRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
	return r._futureInventoryMainOperationQuest
}
