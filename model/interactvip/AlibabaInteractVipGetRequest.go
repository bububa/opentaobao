package interactvip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员淘气值获取 API请求
alibaba.interact.vip.get

提供用户淘气值&用户角色身份查询
*/
type AlibabaInteractVipGetRequest struct {
    model.Params
}

// 初始化AlibabaInteractVipGetRequest对象
func NewAlibabaInteractVipGetRequest() *AlibabaInteractVipGetRequest{
    return &AlibabaInteractVipGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractVipGetRequest) GetApiMethodName() string {
    return "alibaba.interact.vip.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractVipGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
