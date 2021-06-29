package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务类型 API请求
tmall.servicecenter.service.type.queryall

查询天猫服务类型列表
*/
type TmallServicecenterServiceTypeQueryallRequest struct {
    model.Params
}

// 初始化TmallServicecenterServiceTypeQueryallRequest对象
func NewTmallServicecenterServiceTypeQueryallRequest() *TmallServicecenterServiceTypeQueryallRequest{
    return &TmallServicecenterServiceTypeQueryallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServiceTypeQueryallRequest) GetApiMethodName() string {
    return "tmall.servicecenter.service.type.queryall"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServiceTypeQueryallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
