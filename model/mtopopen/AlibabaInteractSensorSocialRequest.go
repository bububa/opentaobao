package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
社交组件 API请求
alibaba.interact.sensor.social

赞，评论 ，关注 新增接口
*/
type AlibabaInteractSensorSocialRequest struct {
    model.Params
    // 系统自动生成
    id   string
}

// 初始化AlibabaInteractSensorSocialRequest对象
func NewAlibabaInteractSensorSocialRequest() *AlibabaInteractSensorSocialRequest{
    return &AlibabaInteractSensorSocialRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorSocialRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.social"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorSocialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorSocialRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlibabaInteractSensorSocialRequest) GetId() string {
    return r.id
}
