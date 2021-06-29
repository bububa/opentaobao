package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-楼层接口pushFloor API请求
alibaba.damai.mev.open.pushfloor

pushFloor
*/
type AlibabaDamaiMevOpenPushfloorRequest struct {
    model.Params
    // 入参pushFloorParam
    pushFloorParam   *ThirdFloorPushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushfloorRequest对象
func NewAlibabaDamaiMevOpenPushfloorRequest() *AlibabaDamaiMevOpenPushfloorRequest{
    return &AlibabaDamaiMevOpenPushfloorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfloorRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushfloor"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfloorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushFloorParam Setter
// 入参pushFloorParam
func (r *AlibabaDamaiMevOpenPushfloorRequest) SetPushFloorParam(pushFloorParam *ThirdFloorPushOpenParam) error {
    r.pushFloorParam = pushFloorParam
    r.Set("push_floor_param", pushFloorParam)
    return nil
}

// PushFloorParam Getter
func (r AlibabaDamaiMevOpenPushfloorRequest) GetPushFloorParam() *ThirdFloorPushOpenParam {
    return r.pushFloorParam
}
