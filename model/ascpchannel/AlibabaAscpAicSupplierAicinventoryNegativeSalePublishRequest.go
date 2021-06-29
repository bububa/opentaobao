package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AIC负卖库存新增和修改接口 APIRequest
alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish

新增负卖库存记录和变更负卖库存记录
*/
type AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest struct {
    model.Params

    // 入参
    futureInventoryMainOperationQuest   *Futureinventorymainoperationquest 

}

func NewAlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest() *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest{
    return &AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) GetApiMethodName() string {
    return "alibaba.ascp.aic.supplier.aicinventory.negative.sale.publish"
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) SetFutureInventoryMainOperationQuest(futureInventoryMainOperationQuest *Futureinventorymainoperationquest) error {
    r.futureInventoryMainOperationQuest = futureInventoryMainOperationQuest
    r.Set("future_inventory_main_operation_quest", futureInventoryMainOperationQuest)
    return nil
}

func (r AlibabaAscpAicSupplierAicinventoryNegativeSalePublishRequest) GetFutureInventoryMainOperationQuest() *Futureinventorymainoperationquest {
    return r.futureInventoryMainOperationQuest
}

