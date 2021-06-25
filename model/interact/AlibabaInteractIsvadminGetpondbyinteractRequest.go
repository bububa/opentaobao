package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据互动实例ID查询奖池信息 APIRequest
alibaba.interact.isvadmin.getpondbyinteract

根据互动实例ID查询奖池信息
*/
type AlibabaInteractIsvadminGetpondbyinteractRequest struct {
    model.Params

    // 互动实例ID
    interactId   string 

}

func NewAlibabaInteractIsvadminGetpondbyinteractRequest() *AlibabaInteractIsvadminGetpondbyinteractRequest{
    return &AlibabaInteractIsvadminGetpondbyinteractRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractIsvadminGetpondbyinteractRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.getpondbyinteract"
}

func (r AlibabaInteractIsvadminGetpondbyinteractRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractIsvadminGetpondbyinteractRequest) SetInteractId(interactId string) error {
    r.interactId = interactId
    r.Set("interact_id", interactId)
    return nil
}

func (r AlibabaInteractIsvadminGetpondbyinteractRequest) GetInteractId() string {
    return r.interactId
}

