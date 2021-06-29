package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
负卖库存失效接口 API请求
alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate

失效负卖库存数据
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest struct {
    model.Params
    // 入参
    futureInventoryMainOperationQuest   *Futureinventorymainoperationquest
}

// 初始化AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest对象
func NewAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest{
    return &AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FutureInventoryMainOperationQuest Setter
// 入参
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) SetFutureInventoryMainOperationQuest(futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
    r.futureInventoryMainOperationQuest = futureInventoryMainOperationQuest
    r.Set("future_inventory_main_operation_quest", futureInventoryMainOperationQuest)
    return nil
}

// FutureInventoryMainOperationQuest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
    return r.futureInventoryMainOperationQuest
}
