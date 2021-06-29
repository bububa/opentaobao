package lsticitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品信息查询 APIRequest
alibaba.lst.ic.item.info.query

查询商品信息
*/
type AlibabaLstIcItemInfoQueryRequest struct {
    model.Params

    // 零售通商品查询参数
    query   *LstItemListParam 

}

func NewAlibabaLstIcItemInfoQueryRequest() *AlibabaLstIcItemInfoQueryRequest{
    return &AlibabaLstIcItemInfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstIcItemInfoQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.ic.item.info.query"
}

func (r AlibabaLstIcItemInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstIcItemInfoQueryRequest) SetQuery(query *LstItemListParam) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaLstIcItemInfoQueryRequest) GetQuery() *LstItemListParam {
    return r.query
}

