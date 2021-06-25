package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV互动应用活动关闭服务 APIRequest
alibaba.interact.activity.unregister

卖家在ISV互动应用中设置的活动主动关闭的服务
*/
type AlibabaInteractActivityUnregisterRequest struct {
    model.Params

    // 互动活动ID
    bizId   string 

}

func NewAlibabaInteractActivityUnregisterRequest() *AlibabaInteractActivityUnregisterRequest{
    return &AlibabaInteractActivityUnregisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractActivityUnregisterRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.unregister"
}

func (r AlibabaInteractActivityUnregisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractActivityUnregisterRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r AlibabaInteractActivityUnregisterRequest) GetBizId() string {
    return r.bizId
}

