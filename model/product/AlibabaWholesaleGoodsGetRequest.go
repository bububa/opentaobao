package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询阿里巴巴批发市场商品详情 API请求
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

// 初始化AlibabaWholesaleGoodsGetRequest对象
func NewAlibabaWholesaleGoodsGetRequest() *AlibabaWholesaleGoodsGetRequest{
    return &AlibabaWholesaleGoodsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleGoodsGetRequest) GetApiMethodName() string {
    return "alibaba.wholesale.goods.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleGoodsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CountryCode Setter
// country_code
func (r *AlibabaWholesaleGoodsGetRequest) SetCountryCode(countryCode string) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

// CountryCode Getter
func (r AlibabaWholesaleGoodsGetRequest) GetCountryCode() string {
    return r.countryCode
}
// Id Setter
// id
func (r *AlibabaWholesaleGoodsGetRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaWholesaleGoodsGetRequest) GetId() string {
    return r.id
}
