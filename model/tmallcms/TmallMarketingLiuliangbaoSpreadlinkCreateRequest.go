package tmallcms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建流量宝活动链接 API请求
tmall.marketing.liuliangbao.spreadlink.create

通过源活动链接和商家ID，创建流量宝活动链接
*/
type TmallMarketingLiuliangbaoSpreadlinkCreateRequest struct {
    model.Params
    // 活动链接，必须为淘系链接
    url   string
}

// 初始化TmallMarketingLiuliangbaoSpreadlinkCreateRequest对象
func NewTmallMarketingLiuliangbaoSpreadlinkCreateRequest() *TmallMarketingLiuliangbaoSpreadlinkCreateRequest{
    return &TmallMarketingLiuliangbaoSpreadlinkCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMarketingLiuliangbaoSpreadlinkCreateRequest) GetApiMethodName() string {
    return "tmall.marketing.liuliangbao.spreadlink.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallMarketingLiuliangbaoSpreadlinkCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Url Setter
// 活动链接，必须为淘系链接
func (r *TmallMarketingLiuliangbaoSpreadlinkCreateRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r TmallMarketingLiuliangbaoSpreadlinkCreateRequest) GetUrl() string {
    return r.url
}
