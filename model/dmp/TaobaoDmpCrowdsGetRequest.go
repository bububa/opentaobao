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
type TaobaoDmpCrowdsGetRequest struct {
    model.Params
}

// 初始化TaobaoDmpCrowdsGetRequest对象
func NewTaobaoDmpCrowdsGetRequest() *TaobaoDmpCrowdsGetRequest{
    return &TaobaoDmpCrowdsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDmpCrowdsGetRequest) GetApiMethodName() string {
    return "taobao.dmp.crowds.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDmpCrowdsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
