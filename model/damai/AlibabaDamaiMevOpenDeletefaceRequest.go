package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面接口deleteFace API请求
alibaba.damai.mev.open.deleteface

deleteFace
*/
type AlibabaDamaiMevOpenDeletefaceRequest struct {
    model.Params
    // 入参deleteFaceParam
    deleteFaceParam   *TicketFaceIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeletefaceRequest对象
func NewAlibabaDamaiMevOpenDeletefaceRequest() *AlibabaDamaiMevOpenDeletefaceRequest{
    return &AlibabaDamaiMevOpenDeletefaceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletefaceRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteface"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletefaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteFaceParam Setter
// 入参deleteFaceParam
func (r *AlibabaDamaiMevOpenDeletefaceRequest) SetDeleteFaceParam(deleteFaceParam *TicketFaceIdOpenParam) error {
    r.deleteFaceParam = deleteFaceParam
    r.Set("delete_face_param", deleteFaceParam)
    return nil
}

// DeleteFaceParam Getter
func (r AlibabaDamaiMevOpenDeletefaceRequest) GetDeleteFaceParam() *TicketFaceIdOpenParam {
    return r.deleteFaceParam
}
