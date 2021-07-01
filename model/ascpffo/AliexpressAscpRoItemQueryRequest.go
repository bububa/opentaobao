package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress退供单明细查询API API请求
aliexpress.ascp.ro.item.query

AE仓发 单个退供单明细查询
*/
type AliexpressAscpRoItemQueryAPIRequest struct {
    model.Params
    // dto
    _returnOrderItemQuery   *ReturnOrderItemQueryDTO
}

// 初始化AliexpressAscpRoItemQueryAPIRequest对象
func NewAliexpressAscpRoItemQueryRequest() *AliexpressAscpRoItemQueryAPIRequest{
    return &AliexpressAscpRoItemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAscpRoItemQueryAPIRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ro.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAscpRoItemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReturnOrderItemQuery Setter
// dto
func (r *AliexpressAscpRoItemQueryAPIRequest) SetReturnOrderItemQuery(_returnOrderItemQuery *ReturnOrderItemQueryDTO) error {
    r._returnOrderItemQuery = _returnOrderItemQuery
    r.Set("return_order_item_query", _returnOrderItemQuery)
    return nil
}

// ReturnOrderItemQuery Getter
func (r AliexpressAscpRoItemQueryAPIRequest) GetReturnOrderItemQuery() *ReturnOrderItemQueryDTO {
    return r._returnOrderItemQuery
}
