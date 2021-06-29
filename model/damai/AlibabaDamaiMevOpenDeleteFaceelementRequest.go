package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement APIRequest
alibaba.damai.mev.open.delete.faceelement

deleteFaceElement
*/
type AlibabaDamaiMevOpenDeleteFaceelementRequest struct {
    model.Params

    // 入参deleteFaceElementParam
    deleteFaceElementParam   *FaceElementIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeleteFaceelementRequest() *AlibabaDamaiMevOpenDeleteFaceelementRequest{
    return &AlibabaDamaiMevOpenDeleteFaceelementRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeleteFaceelementRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.delete.faceelement"
}

func (r AlibabaDamaiMevOpenDeleteFaceelementRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeleteFaceelementRequest) SetDeleteFaceElementParam(deleteFaceElementParam *FaceElementIdOpenParam) error {
    r.deleteFaceElementParam = deleteFaceElementParam
    r.Set("delete_face_element_param", deleteFaceElementParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeleteFaceelementRequest) GetDeleteFaceElementParam() *FaceElementIdOpenParam {
    return r.deleteFaceElementParam
}

