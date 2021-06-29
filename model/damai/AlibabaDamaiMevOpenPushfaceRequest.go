package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面接口pushFace APIRequest
alibaba.damai.mev.open.pushface

pushFace
*/
type AlibabaDamaiMevOpenPushfaceRequest struct {
    model.Params

    // 入参pushFaceParam
    pushFaceParam   *ThirdTicketFacePushOpenParam 

}

func NewAlibabaDamaiMevOpenPushfaceRequest() *AlibabaDamaiMevOpenPushfaceRequest{
    return &AlibabaDamaiMevOpenPushfaceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushfaceRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushface"
}

func (r AlibabaDamaiMevOpenPushfaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushfaceRequest) SetPushFaceParam(pushFaceParam *ThirdTicketFacePushOpenParam) error {
    r.pushFaceParam = pushFaceParam
    r.Set("push_face_param", pushFaceParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushfaceRequest) GetPushFaceParam() *ThirdTicketFacePushOpenParam {
    return r.pushFaceParam
}

