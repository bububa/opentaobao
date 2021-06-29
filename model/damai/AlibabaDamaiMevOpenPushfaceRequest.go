package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面接口pushFace API请求
alibaba.damai.mev.open.pushface

pushFace
*/
type AlibabaDamaiMevOpenPushfaceRequest struct {
    model.Params
    // 入参pushFaceParam
    _pushFaceParam   *ThirdTicketFacePushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushfaceRequest对象
func NewAlibabaDamaiMevOpenPushfaceRequest() *AlibabaDamaiMevOpenPushfaceRequest{
    return &AlibabaDamaiMevOpenPushfaceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfaceRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushface"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushFaceParam Setter
// 入参pushFaceParam
func (r *AlibabaDamaiMevOpenPushfaceRequest) SetPushFaceParam(_pushFaceParam *ThirdTicketFacePushOpenParam) error {
    r._pushFaceParam = _pushFaceParam
    r.Set("push_face_param", _pushFaceParam)
    return nil
}

// PushFaceParam Getter
func (r AlibabaDamaiMevOpenPushfaceRequest) GetPushFaceParam() *ThirdTicketFacePushOpenParam {
    return r._pushFaceParam
}
