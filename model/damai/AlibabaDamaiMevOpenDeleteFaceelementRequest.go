package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement API请求
alibaba.damai.mev.open.delete.faceelement

deleteFaceElement
*/
type AlibabaDamaiMevOpenDeleteFaceelementRequest struct {
    model.Params
    // 入参deleteFaceElementParam
    deleteFaceElementParam   *FaceElementIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeleteFaceelementRequest对象
func NewAlibabaDamaiMevOpenDeleteFaceelementRequest() *AlibabaDamaiMevOpenDeleteFaceelementRequest{
    return &AlibabaDamaiMevOpenDeleteFaceelementRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteFaceelementRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.delete.faceelement"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteFaceelementRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteFaceElementParam Setter
// 入参deleteFaceElementParam
func (r *AlibabaDamaiMevOpenDeleteFaceelementRequest) SetDeleteFaceElementParam(deleteFaceElementParam *FaceElementIdOpenParam) error {
    r.deleteFaceElementParam = deleteFaceElementParam
    r.Set("delete_face_element_param", deleteFaceElementParam)
    return nil
}

// DeleteFaceElementParam Getter
func (r AlibabaDamaiMevOpenDeleteFaceelementRequest) GetDeleteFaceElementParam() *FaceElementIdOpenParam {
    return r.deleteFaceElementParam
}
