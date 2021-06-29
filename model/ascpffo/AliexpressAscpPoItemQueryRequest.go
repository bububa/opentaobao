package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress采购单明细查询API API请求
aliexpress.ascp.po.item.query

AE 供应链仓发 采购单明细查询
*/
type AliexpressAscpPoItemQueryRequest struct {
    model.Params
    // demo
    _purchaseOrderItemQuery   *PurchaseOrderItemQueryDTO
}

// 初始化AliexpressAscpPoItemQueryRequest对象
func NewAliexpressAscpPoItemQueryRequest() *AliexpressAscpPoItemQueryRequest{
    return &AliexpressAscpPoItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpPoItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.po.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpPoItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PurchaseOrderItemQuery Setter
// demo
func (r *AliexpressAscpPoItemQueryRequest) SetPurchaseOrderItemQuery(_purchaseOrderItemQuery *PurchaseOrderItemQueryDTO) error {
    r._purchaseOrderItemQuery = _purchaseOrderItemQuery
    r.Set("purchase_order_item_query", _purchaseOrderItemQuery)
    return nil
}

// PurchaseOrderItemQuery Getter
func (r AliexpressAscpPoItemQueryRequest) GetPurchaseOrderItemQuery() *PurchaseOrderItemQueryDTO {
    return r._purchaseOrderItemQuery
}
