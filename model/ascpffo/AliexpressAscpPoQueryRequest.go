package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress采购单查询API API请求
aliexpress.ascp.po.query

AE仓发业务采购单查询
*/
type AliexpressAscpPoQueryRequest struct {
    model.Params
    // 系统自动生成
    _purchaseOrderQuery   *PurchaseOrderQueryDTO
}

// 初始化AliexpressAscpPoQueryRequest对象
func NewAliexpressAscpPoQueryRequest() *AliexpressAscpPoQueryRequest{
    return &AliexpressAscpPoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpPoQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.po.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpPoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PurchaseOrderQuery Setter
// 系统自动生成
func (r *AliexpressAscpPoQueryRequest) SetPurchaseOrderQuery(_purchaseOrderQuery *PurchaseOrderQueryDTO) error {
    r._purchaseOrderQuery = _purchaseOrderQuery
    r.Set("purchase_order_query", _purchaseOrderQuery)
    return nil
}

// PurchaseOrderQuery Getter
func (r AliexpressAscpPoQueryRequest) GetPurchaseOrderQuery() *PurchaseOrderQueryDTO {
    return r._purchaseOrderQuery
}
