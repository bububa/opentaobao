package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫互动奖池列表 API请求
alibaba.interact.isvadmin.allponds

获取天猫互动奖池列表
*/
type AlibabaInteractIsvadminAllpondsRequest struct {
    model.Params
}

// 初始化AlibabaInteractIsvadminAllpondsRequest对象
func NewAlibabaInteractIsvadminAllpondsRequest() *AlibabaInteractIsvadminAllpondsRequest{
    return &AlibabaInteractIsvadminAllpondsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminAllpondsRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.allponds"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminAllpondsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
