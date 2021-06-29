package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面接口deleteFace APIRequest
alibaba.damai.mev.open.deleteface

deleteFace
*/
type AlibabaDamaiMevOpenDeletefaceRequest struct {
    model.Params

    // 入参deleteFaceParam
    deleteFaceParam   *TicketFaceIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeletefaceRequest() *AlibabaDamaiMevOpenDeletefaceRequest{
    return &AlibabaDamaiMevOpenDeletefaceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeletefaceRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteface"
}

func (r AlibabaDamaiMevOpenDeletefaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeletefaceRequest) SetDeleteFaceParam(deleteFaceParam *TicketFaceIdOpenParam) error {
    r.deleteFaceParam = deleteFaceParam
    r.Set("delete_face_param", deleteFaceParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeletefaceRequest) GetDeleteFaceParam() *TicketFaceIdOpenParam {
    return r.deleteFaceParam
}

