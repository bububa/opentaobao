package tuike

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广商品查询接口 API请求
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

// 初始化AlibabaTuikeOfferGetRequest对象
func NewAlibabaTuikeOfferGetRequest() *AlibabaTuikeOfferGetRequest{
    return &AlibabaTuikeOfferGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTuikeOfferGetRequest) GetApiMethodName() string {
    return "alibaba.tuike.offer.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTuikeOfferGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvCode Setter
// 标识调用方
func (r *AlibabaTuikeOfferGetRequest) SetIsvCode(isvCode string) error {
    r.isvCode = isvCode
    r.Set("isv_code", isvCode)
    return nil
}

// IsvCode Getter
func (r AlibabaTuikeOfferGetRequest) GetIsvCode() string {
    return r.isvCode
}
// QueryString Setter
// 搜索查询参数(json)
func (r *AlibabaTuikeOfferGetRequest) SetQueryString(queryString string) error {
    r.queryString = queryString
    r.Set("query_string", queryString)
    return nil
}

// QueryString Getter
func (r AlibabaTuikeOfferGetRequest) GetQueryString() string {
    return r.queryString
}
