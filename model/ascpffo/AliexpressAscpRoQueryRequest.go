package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress退供单查询API APIRequest
aliexpress.ascp.ro.query

AE仓发商家单个退供单查询接口
*/
type AliexpressAscpRoQueryRequest struct {
    model.Params

    // dto
    returnOrderQuery   *ReturnOrderQueryDto 

}

func NewAliexpressAscpRoQueryRequest() *AliexpressAscpRoQueryRequest{
    return &AliexpressAscpRoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpRoQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.ro.query"
}

func (r AliexpressAscpRoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpRoQueryRequest) SetReturnOrderQuery(returnOrderQuery *ReturnOrderQueryDto) error {
    r.returnOrderQuery = returnOrderQuery
    r.Set("return_order_query", returnOrderQuery)
    return nil
}

func (r AliexpressAscpRoQueryRequest) GetReturnOrderQuery() *ReturnOrderQueryDto {
    return r.returnOrderQuery
}

