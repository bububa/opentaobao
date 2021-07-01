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
type AlibabaWholesaleGoodsGetAPIRequest struct {
    model.Params
    // country_code
    _countryCode   string
    // id
    _id   string
}

// 初始化AlibabaWholesaleGoodsGetAPIRequest对象
func NewAlibabaWholesaleGoodsGetRequest() *AlibabaWholesaleGoodsGetAPIRequest{
    return &AlibabaWholesaleGoodsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWholesaleGoodsGetAPIRequest) GetApiMethodName() string {
    return "alibaba.wholesale.goods.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWholesaleGoodsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CountryCode Setter
// country_code
func (r *AlibabaWholesaleGoodsGetAPIRequest) SetCountryCode(_countryCode string) error {
    r._countryCode = _countryCode
    r.Set("country_code", _countryCode)
    return nil
}

// CountryCode Getter
func (r AlibabaWholesaleGoodsGetAPIRequest) GetCountryCode() string {
    return r._countryCode
}
// Id Setter
// id
func (r *AlibabaWholesaleGoodsGetAPIRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaWholesaleGoodsGetAPIRequest) GetId() string {
    return r._id
}
