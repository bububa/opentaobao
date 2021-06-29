package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress采购单查询API APIRequest
aliexpress.ascp.po.query

AE仓发业务采购单查询
*/
type AliexpressAscpPoQueryRequest struct {
    model.Params

    // 系统自动生成
    purchaseOrderQuery   *PurchaseOrderQueryDto 

}

func NewAliexpressAscpPoQueryRequest() *AliexpressAscpPoQueryRequest{
    return &AliexpressAscpPoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpPoQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.po.query"
}

func (r AliexpressAscpPoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpPoQueryRequest) SetPurchaseOrderQuery(purchaseOrderQuery *PurchaseOrderQueryDto) error {
    r.purchaseOrderQuery = purchaseOrderQuery
    r.Set("purchase_order_query", purchaseOrderQuery)
    return nil
}

func (r AliexpressAscpPoQueryRequest) GetPurchaseOrderQuery() *PurchaseOrderQueryDto {
    return r.purchaseOrderQuery
}

