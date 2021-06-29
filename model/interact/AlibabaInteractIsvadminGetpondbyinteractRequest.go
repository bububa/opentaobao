package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据互动实例ID查询奖池信息 API请求
alibaba.interact.isvadmin.getpondbyinteract

根据互动实例ID查询奖池信息
*/
type AlibabaInteractIsvadminGetpondbyinteractRequest struct {
    model.Params
    // 互动实例ID
    interactId   string
}

// 初始化AlibabaInteractIsvadminGetpondbyinteractRequest对象
func NewAlibabaInteractIsvadminGetpondbyinteractRequest() *AlibabaInteractIsvadminGetpondbyinteractRequest{
    return &AlibabaInteractIsvadminGetpondbyinteractRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminGetpondbyinteractRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.getpondbyinteract"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminGetpondbyinteractRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InteractId Setter
// 互动实例ID
func (r *AlibabaInteractIsvadminGetpondbyinteractRequest) SetInteractId(interactId string) error {
    r.interactId = interactId
    r.Set("interact_id", interactId)
    return nil
}

// InteractId Getter
func (r AlibabaInteractIsvadminGetpondbyinteractRequest) GetInteractId() string {
    return r.interactId
}
