package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
社交组件 APIRequest
alibaba.interact.sensor.social

赞，评论 ，关注 新增接口
*/
type AlibabaInteractSensorSocialRequest struct {
    model.Params

    // 系统自动生成
    id   string 

}

func NewAlibabaInteractSensorSocialRequest() *AlibabaInteractSensorSocialRequest{
    return &AlibabaInteractSensorSocialRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSensorSocialRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.social"
}

func (r AlibabaInteractSensorSocialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractSensorSocialRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaInteractSensorSocialRequest) GetId() string {
    return r.id
}

