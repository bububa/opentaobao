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
type AlibabaInteractIsvadminGetpondbyinteractAPIRequest struct {
    model.Params
    // 互动实例ID
    _interactId   string
}

// 初始化AlibabaInteractIsvadminGetpondbyinteractAPIRequest对象
func NewAlibabaInteractIsvadminGetpondbyinteractRequest() *AlibabaInteractIsvadminGetpondbyinteractAPIRequest{
    return &AlibabaInteractIsvadminGetpondbyinteractAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminGetpondbyinteractAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.getpondbyinteract"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminGetpondbyinteractAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InteractId Setter
// 互动实例ID
func (r *AlibabaInteractIsvadminGetpondbyinteractAPIRequest) SetInteractId(_interactId string) error {
    r._interactId = _interactId
    r.Set("interact_id", _interactId)
    return nil
}

// InteractId Getter
func (r AlibabaInteractIsvadminGetpondbyinteractAPIRequest) GetInteractId() string {
    return r._interactId
}
