package dmp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询人群服务 API请求
taobao.dmp.crowds.get

查询人群服务
*/
type TaobaoDmpCrowdsGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoDmpCrowdsGetAPIRequest对象
func NewTaobaoDmpCrowdsGetRequest() *TaobaoDmpCrowdsGetAPIRequest{
    return &TaobaoDmpCrowdsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDmpCrowdsGetAPIRequest) GetApiMethodName() string {
    return "taobao.dmp.crowds.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDmpCrowdsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
