package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小铺isv推广流量活动到流量聚乐部 API请求
alibaba.interact.activity.pushtoalicom

涉及到流量包的小铺isv，将活动推送到流量聚乐部
*/
type AlibabaInteractActivityPushtoalicomRequest struct {
    model.Params
    // 推送到流量聚乐部的banner图
    _bannerUrl   string
    // 推送到流量聚乐部的活动bizid
    _bizId   string
}

// 初始化AlibabaInteractActivityPushtoalicomRequest对象
func NewAlibabaInteractActivityPushtoalicomRequest() *AlibabaInteractActivityPushtoalicomRequest{
    return &AlibabaInteractActivityPushtoalicomRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityPushtoalicomRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.pushtoalicom"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityPushtoalicomRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BannerUrl Setter
// 推送到流量聚乐部的banner图
func (r *AlibabaInteractActivityPushtoalicomRequest) SetBannerUrl(_bannerUrl string) error {
    r._bannerUrl = _bannerUrl
    r.Set("banner_url", _bannerUrl)
    return nil
}

// BannerUrl Getter
func (r AlibabaInteractActivityPushtoalicomRequest) GetBannerUrl() string {
    return r._bannerUrl
}
// BizId Setter
// 推送到流量聚乐部的活动bizid
func (r *AlibabaInteractActivityPushtoalicomRequest) SetBizId(_bizId string) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r AlibabaInteractActivityPushtoalicomRequest) GetBizId() string {
    return r._bizId
}
