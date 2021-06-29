package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广链接 API请求
aliexpress.social.item.promotion

获取商品社交推广链接
*/
type AliexpressSocialItemPromotionRequest struct {
    model.Params
    // 推广的商品链接
    _targetUrl   string
    // 子渠道号
    _af   string
    // campaign Id
    _cn   string
    // creative id
    _cv   string
    // click id
    _dp   string
}

// 初始化AliexpressSocialItemPromotionRequest对象
func NewAliexpressSocialItemPromotionRequest() *AliexpressSocialItemPromotionRequest{
    return &AliexpressSocialItemPromotionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialItemPromotionRequest) GetApiMethodName() string {
    return "aliexpress.social.item.promotion"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialItemPromotionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TargetUrl Setter
// 推广的商品链接
func (r *AliexpressSocialItemPromotionRequest) SetTargetUrl(_targetUrl string) error {
    r._targetUrl = _targetUrl
    r.Set("target_url", _targetUrl)
    return nil
}

// TargetUrl Getter
func (r AliexpressSocialItemPromotionRequest) GetTargetUrl() string {
    return r._targetUrl
}
// Af Setter
// 子渠道号
func (r *AliexpressSocialItemPromotionRequest) SetAf(_af string) error {
    r._af = _af
    r.Set("af", _af)
    return nil
}

// Af Getter
func (r AliexpressSocialItemPromotionRequest) GetAf() string {
    return r._af
}
// Cn Setter
// campaign Id
func (r *AliexpressSocialItemPromotionRequest) SetCn(_cn string) error {
    r._cn = _cn
    r.Set("cn", _cn)
    return nil
}

// Cn Getter
func (r AliexpressSocialItemPromotionRequest) GetCn() string {
    return r._cn
}
// Cv Setter
// creative id
func (r *AliexpressSocialItemPromotionRequest) SetCv(_cv string) error {
    r._cv = _cv
    r.Set("cv", _cv)
    return nil
}

// Cv Getter
func (r AliexpressSocialItemPromotionRequest) GetCv() string {
    return r._cv
}
// Dp Setter
// click id
func (r *AliexpressSocialItemPromotionRequest) SetDp(_dp string) error {
    r._dp = _dp
    r.Set("dp", _dp)
    return nil
}

// Dp Getter
func (r AliexpressSocialItemPromotionRequest) GetDp() string {
    return r._dp
}
