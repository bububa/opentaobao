package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress退供单明细查询API APIRequest
aliexpress.ascp.ro.item.query

AE仓发 单个退供单明细查询
*/
type AliexpressAscpRoItemQueryRequest struct {
    model.Params

    // dto
    returnOrderItemQuery   *ReturnOrderItemQueryDto 

}

func NewAliexpressAscpRoItemQueryRequest() *AliexpressAscpRoItemQueryRequest{
    return &AliexpressAscpRoItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpRoItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ro.item.query"
}

func (r AliexpressAscpRoItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpRoItemQueryRequest) SetReturnOrderItemQuery(returnOrderItemQuery *ReturnOrderItemQueryDto) error {
    r.returnOrderItemQuery = returnOrderItemQuery
    r.Set("return_order_item_query", returnOrderItemQuery)
    return nil
}

func (r AliexpressAscpRoItemQueryRequest) GetReturnOrderItemQuery() *ReturnOrderItemQueryDto {
    return r.returnOrderItemQuery
}

