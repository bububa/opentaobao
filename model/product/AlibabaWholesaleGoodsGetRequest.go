package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询阿里巴巴批发市场商品详情 APIRequest
alibaba.wholesale.goods.get

查询阿里巴巴批发市场商品详情
*/
type AlibabaWholesaleGoodsGetRequest struct {
    model.Params

    // country_code
    countryCode   string 

    // id
    id   string 

}

func NewAlibabaWholesaleGoodsGetRequest() *AlibabaWholesaleGoodsGetRequest{
    return &AlibabaWholesaleGoodsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWholesaleGoodsGetRequest) GetApiMethodName() string {
    return "alibaba.wholesale.goods.get"
}

func (r AlibabaWholesaleGoodsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWholesaleGoodsGetRequest) SetCountryCode(countryCode string) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

func (r AlibabaWholesaleGoodsGetRequest) GetCountryCode() string {
    return r.countryCode
}

func (r *AlibabaWholesaleGoodsGetRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaWholesaleGoodsGetRequest) GetId() string {
    return r.id
}

