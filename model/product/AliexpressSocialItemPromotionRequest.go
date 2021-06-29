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
    targetUrl   string
    // 子渠道号
    af   string
    // campaign Id
    cn   string
    // creative id
    cv   string
    // click id
    dp   string
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
func (r *AliexpressSocialItemPromotionRequest) SetTargetUrl(targetUrl string) error {
    r.targetUrl = targetUrl
    r.Set("target_url", targetUrl)
    return nil
}

// TargetUrl Getter
func (r AliexpressSocialItemPromotionRequest) GetTargetUrl() string {
    return r.targetUrl
}
// Af Setter
// 子渠道号
func (r *AliexpressSocialItemPromotionRequest) SetAf(af string) error {
    r.af = af
    r.Set("af", af)
    return nil
}

// Af Getter
func (r AliexpressSocialItemPromotionRequest) GetAf() string {
    return r.af
}
// Cn Setter
// campaign Id
func (r *AliexpressSocialItemPromotionRequest) SetCn(cn string) error {
    r.cn = cn
    r.Set("cn", cn)
    return nil
}

// Cn Getter
func (r AliexpressSocialItemPromotionRequest) GetCn() string {
    return r.cn
}
// Cv Setter
// creative id
func (r *AliexpressSocialItemPromotionRequest) SetCv(cv string) error {
    r.cv = cv
    r.Set("cv", cv)
    return nil
}

// Cv Getter
func (r AliexpressSocialItemPromotionRequest) GetCv() string {
    return r.cv
}
// Dp Setter
// click id
func (r *AliexpressSocialItemPromotionRequest) SetDp(dp string) error {
    r.dp = dp
    r.Set("dp", dp)
    return nil
}

// Dp Getter
func (r AliexpressSocialItemPromotionRequest) GetDp() string {
    return r.dp
}
