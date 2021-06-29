package tuike

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广商品查询接口2.0 API请求
alibaba.tuike.offer.get.pro

查询1688推客平台卖家推广中的商品信息
*/
type AlibabaTuikeOfferGetProRequest struct {
    model.Params
    // 用户ID
    _loginId   string
    // 标识调用方
    _isvCode   string
    // 搜索查询参数(json)
    _queryString   string
}

// 初始化AlibabaTuikeOfferGetProRequest对象
func NewAlibabaTuikeOfferGetProRequest() *AlibabaTuikeOfferGetProRequest{
    return &AlibabaTuikeOfferGetProRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTuikeOfferGetProRequest) GetApiMethodName() string {
    return "alibaba.tuike.offer.get.pro"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTuikeOfferGetProRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LoginId Setter
// 用户ID
func (r *AlibabaTuikeOfferGetProRequest) SetLoginId(_loginId string) error {
    r._loginId = _loginId
    r.Set("login_id", _loginId)
    return nil
}

// LoginId Getter
func (r AlibabaTuikeOfferGetProRequest) GetLoginId() string {
    return r._loginId
}
// IsvCode Setter
// 标识调用方
func (r *AlibabaTuikeOfferGetProRequest) SetIsvCode(_isvCode string) error {
    r._isvCode = _isvCode
    r.Set("isv_code", _isvCode)
    return nil
}

// IsvCode Getter
func (r AlibabaTuikeOfferGetProRequest) GetIsvCode() string {
    return r._isvCode
}
// QueryString Setter
// 搜索查询参数(json)
func (r *AlibabaTuikeOfferGetProRequest) SetQueryString(_queryString string) error {
    r._queryString = _queryString
    r.Set("query_string", _queryString)
    return nil
}

// QueryString Getter
func (r AlibabaTuikeOfferGetProRequest) GetQueryString() string {
    return r._queryString
}
