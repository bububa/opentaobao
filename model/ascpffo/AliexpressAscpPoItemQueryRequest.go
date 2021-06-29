package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress采购单明细查询API APIRequest
aliexpress.ascp.po.item.query

AE 供应链仓发 采购单明细查询
*/
type AliexpressAscpPoItemQueryRequest struct {
    model.Params

    // demo
    purchaseOrderItemQuery   *PurchaseOrderItemQueryDto 

}

func NewAliexpressAscpPoItemQueryRequest() *AliexpressAscpPoItemQueryRequest{
    return &AliexpressAscpPoItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpPoItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.po.item.query"
}

func (r AliexpressAscpPoItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpPoItemQueryRequest) SetPurchaseOrderItemQuery(purchaseOrderItemQuery *PurchaseOrderItemQueryDto) error {
    r.purchaseOrderItemQuery = purchaseOrderItemQuery
    r.Set("purchase_order_item_query", purchaseOrderItemQuery)
    return nil
}

func (r AliexpressAscpPoItemQueryRequest) GetPurchaseOrderItemQuery() *PurchaseOrderItemQueryDto {
    return r.purchaseOrderItemQuery
}

