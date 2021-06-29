package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票面元素接口pushFaceElement APIRequest
alibaba.damai.mev.open.pushfaceelement

pushFaceElement
*/
type AlibabaDamaiMevOpenPushfaceelementRequest struct {
    model.Params

    // 入参pushFaceElementParamList
    pushFaceElementParamList   []ThirdFaceElementPushOpenParam 

}

func NewAlibabaDamaiMevOpenPushfaceelementRequest() *AlibabaDamaiMevOpenPushfaceelementRequest{
    return &AlibabaDamaiMevOpenPushfaceelementRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushfaceelementRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushfaceelement"
}

func (r AlibabaDamaiMevOpenPushfaceelementRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushfaceelementRequest) SetPushFaceElementParamList(pushFaceElementParamList []ThirdFaceElementPushOpenParam) error {
    r.pushFaceElementParamList = pushFaceElementParamList
    r.Set("push_face_element_param_list", pushFaceElementParamList)
    return nil
}

func (r AlibabaDamaiMevOpenPushfaceelementRequest) GetPushFaceElementParamList() []ThirdFaceElementPushOpenParam {
    return r.pushFaceElementParamList
}

