package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest AIC负卖库存新增和修改接口 API请求
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish
//
// 新增负卖库存记录和变更负卖库存记录
type AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest struct {
	model.Params
	// 入参
	_futureInventoryMainOperationQuest *Futureinventorymainoperationquest
}

// NewAlibabaascpaicsupplieraicinventorynegativesalepublishRequest 初始化AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest对象
func NewAlibabaascpaicsupplieraicinventorynegativesalepublishRequest() *AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest {
	return &AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureInventoryMainOperationQuest is FutureInventoryMainOperationQuest Setter
// 入参
func (r *AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest) SetFutureInventoryMainOperationQuest(_futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
	r._futureInventoryMainOperationQuest = _futureInventoryMainOperationQuest
	r.Set("future_inventory_main_operation_quest", _futureInventoryMainOperationQuest)
	return nil
}

// GetFutureInventoryMainOperationQuest FutureInventoryMainOperationQuest Getter
func (r AlibabaascpaicsupplieraicinventorynegativesalepublishAPIRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
	return r._futureInventoryMainOperationQuest
}
