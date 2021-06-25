package tmallcms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建流量宝活动链接 APIRequest
tmall.marketing.liuliangbao.spreadlink.create

通过源活动链接和商家ID，创建流量宝活动链接
*/
type TmallMarketingLiuliangbaoSpreadlinkCreateRequest struct {
    model.Params

    // 活动链接，必须为淘系链接
    url   string 

}

func NewTmallMarketingLiuliangbaoSpreadlinkCreateRequest() *TmallMarketingLiuliangbaoSpreadlinkCreateRequest{
    return &TmallMarketingLiuliangbaoSpreadlinkCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMarketingLiuliangbaoSpreadlinkCreateRequest) GetApiMethodName() string {
    return "tmall.marketing.liuliangbao.spreadlink.create"
}

func (r TmallMarketingLiuliangbaoSpreadlinkCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMarketingLiuliangbaoSpreadlinkCreateRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TmallMarketingLiuliangbaoSpreadlinkCreateRequest) GetUrl() string {
    return r.url
}

