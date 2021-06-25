package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广链接 APIRequest
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

func NewAliexpressSocialItemPromotionRequest() *AliexpressSocialItemPromotionRequest{
    return &AliexpressSocialItemPromotionRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialItemPromotionRequest) GetApiMethodName() string {
    return "aliexpress.social.item.promotion"
}

func (r AliexpressSocialItemPromotionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSocialItemPromotionRequest) SetTargetUrl(targetUrl string) error {
    r.targetUrl = targetUrl
    r.Set("target_url", targetUrl)
    return nil
}

func (r AliexpressSocialItemPromotionRequest) GetTargetUrl() string {
    return r.targetUrl
}

func (r *AliexpressSocialItemPromotionRequest) SetAf(af string) error {
    r.af = af
    r.Set("af", af)
    return nil
}

func (r AliexpressSocialItemPromotionRequest) GetAf() string {
    return r.af
}

func (r *AliexpressSocialItemPromotionRequest) SetCn(cn string) error {
    r.cn = cn
    r.Set("cn", cn)
    return nil
}

func (r AliexpressSocialItemPromotionRequest) GetCn() string {
    return r.cn
}

func (r *AliexpressSocialItemPromotionRequest) SetCv(cv string) error {
    r.cv = cv
    r.Set("cv", cv)
    return nil
}

func (r AliexpressSocialItemPromotionRequest) GetCv() string {
    return r.cv
}

func (r *AliexpressSocialItemPromotionRequest) SetDp(dp string) error {
    r.dp = dp
    r.Set("dp", dp)
    return nil
}

func (r AliexpressSocialItemPromotionRequest) GetDp() string {
    return r.dp
}

