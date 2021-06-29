package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店查询接口 API请求
alibaba.wdk.shop.query

根据门店code查询门店信息
*/
type AlibabaWdkShopQueryRequest struct {
    model.Params
    // 如果不传，返回所有
    _ouCode   string
}

// 初始化AlibabaWdkShopQueryRequest对象
func NewAlibabaWdkShopQueryRequest() *AlibabaWdkShopQueryRequest{
    return &AlibabaWdkShopQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkShopQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.shop.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkShopQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuCode Setter
// 如果不传，返回所有
func (r *AlibabaWdkShopQueryRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaWdkShopQueryRequest) GetOuCode() string {
    return r._ouCode
}
