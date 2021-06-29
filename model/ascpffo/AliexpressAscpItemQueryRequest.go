package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AliExpress货品查询查询API APIRequest
aliexpress.ascp.item.query

AE货品查询API
*/
type AliexpressAscpItemQueryRequest struct {
    model.Params

    // DTO
    scItemQuery   *ScItemQueryDto 

}

func NewAliexpressAscpItemQueryRequest() *AliexpressAscpItemQueryRequest{
    return &AliexpressAscpItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAscpItemQueryRequest) GetApiMethodName() string {
    return "aliexpress.ascp.item.query"
}

func (r AliexpressAscpItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAscpItemQueryRequest) SetScItemQuery(scItemQuery *ScItemQueryDto) error {
    r.scItemQuery = scItemQuery
    r.Set("sc_item_query", scItemQuery)
    return nil
}

func (r AliexpressAscpItemQueryRequest) GetScItemQuery() *ScItemQueryDto {
    return r.scItemQuery
}

