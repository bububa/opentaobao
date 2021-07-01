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
type TmallServicecenterServiceTypeQueryallAPIRequest struct {
    model.Params
}

// 初始化TmallServicecenterServiceTypeQueryallAPIRequest对象
func NewTmallServicecenterServiceTypeQueryallRequest() *TmallServicecenterServiceTypeQueryallAPIRequest{
    return &TmallServicecenterServiceTypeQueryallAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServiceTypeQueryallAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.service.type.queryall"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServiceTypeQueryallAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
