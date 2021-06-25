package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小铺isv推广流量活动到流量聚乐部 APIRequest
alibaba.interact.activity.pushtoalicom

涉及到流量包的小铺isv，将活动推送到流量聚乐部
*/
type AlibabaInteractActivityPushtoalicomRequest struct {
    model.Params

    // 推送到流量聚乐部的banner图
    bannerUrl   string 

    // 推送到流量聚乐部的活动bizid
    bizId   string 

}

func NewAlibabaInteractActivityPushtoalicomRequest() *AlibabaInteractActivityPushtoalicomRequest{
    return &AlibabaInteractActivityPushtoalicomRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractActivityPushtoalicomRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.pushtoalicom"
}

func (r AlibabaInteractActivityPushtoalicomRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractActivityPushtoalicomRequest) SetBannerUrl(bannerUrl string) error {
    r.bannerUrl = bannerUrl
    r.Set("banner_url", bannerUrl)
    return nil
}

func (r AlibabaInteractActivityPushtoalicomRequest) GetBannerUrl() string {
    return r.bannerUrl
}

func (r *AlibabaInteractActivityPushtoalicomRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r AlibabaInteractActivityPushtoalicomRequest) GetBizId() string {
    return r.bizId
}

