package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
负卖库存失效接口 APIRequest
alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate

失效负卖库存数据
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest struct {
    model.Params

    // 入参
    futureInventoryMainOperationQuest   *Futureinventorymainoperationquest 

}

func NewAlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest{
    return &AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.invalidate"
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) SetFutureInventoryMainOperationQuest(futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
    r.futureInventoryMainOperationQuest = futureInventoryMainOperationQuest
    r.Set("future_inventory_main_operation_quest", futureInventoryMainOperationQuest)
    return nil
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSaleInvalidateRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
    return r.futureInventoryMainOperationQuest
}

