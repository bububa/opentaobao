package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-楼层接口pushFloor APIRequest
alibaba.damai.mev.open.pushfloor

pushFloor
*/
type AlibabaDamaiMevOpenPushfloorRequest struct {
    model.Params

    // 入参pushFloorParam
    pushFloorParam   *ThirdFloorPushOpenParam 

}

func NewAlibabaDamaiMevOpenPushfloorRequest() *AlibabaDamaiMevOpenPushfloorRequest{
    return &AlibabaDamaiMevOpenPushfloorRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushfloorRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushfloor"
}

func (r AlibabaDamaiMevOpenPushfloorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushfloorRequest) SetPushFloorParam(pushFloorParam *ThirdFloorPushOpenParam) error {
    r.pushFloorParam = pushFloorParam
    r.Set("push_floor_param", pushFloorParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushfloorRequest) GetPushFloorParam() *ThirdFloorPushOpenParam {
    return r.pushFloorParam
}

