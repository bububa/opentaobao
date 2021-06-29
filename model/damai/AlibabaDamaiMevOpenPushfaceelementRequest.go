package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面元素接口pushFaceElement API请求
alibaba.damai.mev.open.pushfaceelement

pushFaceElement
*/
type AlibabaDamaiMevOpenPushfaceelementRequest struct {
    model.Params
    // 入参pushFaceElementParamList
    pushFaceElementParamList   []ThirdFaceElementPushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushfaceelementRequest对象
func NewAlibabaDamaiMevOpenPushfaceelementRequest() *AlibabaDamaiMevOpenPushfaceelementRequest{
    return &AlibabaDamaiMevOpenPushfaceelementRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushfaceelementRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushfaceelement"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushfaceelementRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushFaceElementParamList Setter
// 入参pushFaceElementParamList
func (r *AlibabaDamaiMevOpenPushfaceelementRequest) SetPushFaceElementParamList(pushFaceElementParamList []ThirdFaceElementPushOpenParam) error {
    r.pushFaceElementParamList = pushFaceElementParamList
    r.Set("push_face_element_param_list", pushFaceElementParamList)
    return nil
}

// PushFaceElementParamList Getter
func (r AlibabaDamaiMevOpenPushfaceelementRequest) GetPushFaceElementParamList() []ThirdFaceElementPushOpenParam {
    return r.pushFaceElementParamList
}
