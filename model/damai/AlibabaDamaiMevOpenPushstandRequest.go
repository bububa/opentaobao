package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-看台接口pushStand API请求
alibaba.damai.mev.open.pushstand

pushStand
*/
type AlibabaDamaiMevOpenPushstandRequest struct {
    model.Params
    // 入参pushStandParam
    _pushStandParam   *ThirdStandPushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushstandRequest对象
func NewAlibabaDamaiMevOpenPushstandRequest() *AlibabaDamaiMevOpenPushstandRequest{
    return &AlibabaDamaiMevOpenPushstandRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushstandRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushstand"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushstandRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushStandParam Setter
// 入参pushStandParam
func (r *AlibabaDamaiMevOpenPushstandRequest) SetPushStandParam(_pushStandParam *ThirdStandPushOpenParam) error {
    r._pushStandParam = _pushStandParam
    r.Set("push_stand_param", _pushStandParam)
    return nil
}

// PushStandParam Getter
func (r AlibabaDamaiMevOpenPushstandRequest) GetPushStandParam() *ThirdStandPushOpenParam {
    return r._pushStandParam
}
