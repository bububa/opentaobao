package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest 负卖库存失效接口 API请求
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate
//
// 失效负卖库存数据
type AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest struct {
	model.Params
	// 入参
	_futureInventoryMainOperationQuest *Futureinventorymainoperationquest
}

// NewAlibabaascpaicsupplieraicinventorynegativesaleinvalidateRequest 初始化AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest对象
func NewAlibabaascpaicsupplieraicinventorynegativesaleinvalidateRequest() *AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest {
	return &AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureInventoryMainOperationQuest is FutureInventoryMainOperationQuest Setter
// 入参
func (r *AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest) SetFutureInventoryMainOperationQuest(_futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
	r._futureInventoryMainOperationQuest = _futureInventoryMainOperationQuest
	r.Set("future_inventory_main_operation_quest", _futureInventoryMainOperationQuest)
	return nil
}

// GetFutureInventoryMainOperationQuest FutureInventoryMainOperationQuest Getter
func (r AlibabaascpaicsupplieraicinventorynegativesaleinvalidateAPIRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
	return r._futureInventoryMainOperationQuest
}
