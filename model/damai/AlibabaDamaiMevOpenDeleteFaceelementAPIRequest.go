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
type AlibabaDamaiMevOpenDeleteFaceelementAPIRequest struct {
    model.Params
    // 入参deleteFaceElementParam
    _deleteFaceElementParam   *FaceElementIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeleteFaceelementAPIRequest对象
func NewAlibabaDamaiMevOpenDeleteFaceelementRequest() *AlibabaDamaiMevOpenDeleteFaceelementAPIRequest{
    return &AlibabaDamaiMevOpenDeleteFaceelementAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.delete.faceelement"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteFaceElementParam Setter
// 入参deleteFaceElementParam
func (r *AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) SetDeleteFaceElementParam(_deleteFaceElementParam *FaceElementIdOpenParam) error {
    r._deleteFaceElementParam = _deleteFaceElementParam
    r.Set("delete_face_element_param", _deleteFaceElementParam)
    return nil
}

// DeleteFaceElementParam Getter
func (r AlibabaDamaiMevOpenDeleteFaceelementAPIRequest) GetDeleteFaceElementParam() *FaceElementIdOpenParam {
    return r._deleteFaceElementParam
}
