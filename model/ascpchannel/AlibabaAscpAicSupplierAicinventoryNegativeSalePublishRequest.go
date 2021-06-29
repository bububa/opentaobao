package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AIC负卖库存新增和修改接口 API请求
alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish

新增负卖库存记录和变更负卖库存记录
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest struct {
    model.Params
    // 入参
    futureInventoryMainOperationQuest   *Futureinventorymainoperationquest
}

// 初始化AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest对象
func NewAlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest{
    return &AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FutureInventoryMainOperationQuest Setter
// 入参
func (r *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) SetFutureInventoryMainOperationQuest(futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
    r.futureInventoryMainOperationQuest = futureInventoryMainOperationQuest
    r.Set("future_inventory_main_operation_quest", futureInventoryMainOperationQuest)
    return nil
}

// FutureInventoryMainOperationQuest Getter
func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
    return r.futureInventoryMainOperationQuest
}
