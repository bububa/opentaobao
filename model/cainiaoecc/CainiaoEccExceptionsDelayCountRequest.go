package cainiaoecc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟控制塔包裹滞留异常统计信息获取 API请求
cainiao.ecc.exceptions.delay.count

菜鸟控制塔包裹滞留异常统计信息获取
*/
type CainiaoEccExceptionsDelayCountRequest struct {
    model.Params
}

// 初始化CainiaoEccExceptionsDelayCountRequest对象
func NewCainiaoEccExceptionsDelayCountRequest() *CainiaoEccExceptionsDelayCountRequest{
    return &CainiaoEccExceptionsDelayCountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEccExceptionsDelayCountRequest) GetApiMethodName() string {
    return "cainiao.ecc.exceptions.delay.count"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEccExceptionsDelayCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
