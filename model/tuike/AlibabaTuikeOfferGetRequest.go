package tuike

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广商品查询接口 APIRequest
alibaba.tuike.offer.get

查询1688推客平台卖家推广中的商品信息
*/
type AlibabaTuikeOfferGetRequest struct {
    model.Params

    // 标识调用方
    isvCode   string 

    // 搜索查询参数(json)
    queryString   string 

}

func NewAlibabaTuikeOfferGetRequest() *AlibabaTuikeOfferGetRequest{
    return &AlibabaTuikeOfferGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTuikeOfferGetRequest) GetApiMethodName() string {
    return "alibaba.tuike.offer.get"
}

func (r AlibabaTuikeOfferGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTuikeOfferGetRequest) SetIsvCode(isvCode string) error {
    r.isvCode = isvCode
    r.Set("isv_code", isvCode)
    return nil
}

func (r AlibabaTuikeOfferGetRequest) GetIsvCode() string {
    return r.isvCode
}

func (r *AlibabaTuikeOfferGetRequest) SetQueryString(queryString string) error {
    r.queryString = queryString
    r.Set("query_string", queryString)
    return nil
}

func (r AlibabaTuikeOfferGetRequest) GetQueryString() string {
    return r.queryString
}

