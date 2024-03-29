package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest AIC负卖库存新增和修改接口 API请求
// alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish
//
// 新增负卖库存记录和变更负卖库存记录
type AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest struct {
	model.Params
	// 入参
	_futureInventoryMainOperationQuest *Futureinventorymainoperationquest
}

// NewAlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest 初始化AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest对象
func NewAlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest {
	return &AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest) Reset() {
	r._futureInventoryMainOperationQuest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFutureInventoryMainOperationQuest is FutureInventoryMainOperationQuest Setter
// 入参
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest) SetFutureInventoryMainOperationQuest(_futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
	r._futureInventoryMainOperationQuest = _futureInventoryMainOperationQuest
	r.Set("future_inventory_main_operation_quest", _futureInventoryMainOperationQuest)
	return nil
}

// GetFutureInventoryMainOperationQuest FutureInventoryMainOperationQuest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
	return r._futureInventoryMainOperationQuest
}

var poolAlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest()
	},
}

// GetAlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest 从 sync.Pool 获取 AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest
func GetAlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest {
	return poolAlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest.Get().(*AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest 将 AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest(v *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest) {
	v.Reset()
	poolAlibabaAscpAicSupplierAicinventoryNegativeSalePublishAPIRequest.Put(v)
}
