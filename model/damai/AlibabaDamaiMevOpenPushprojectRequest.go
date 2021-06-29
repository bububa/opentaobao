package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-项目接口pushProject API请求
alibaba.damai.mev.open.pushproject

pushProject
*/
type AlibabaDamaiMevOpenPushprojectRequest struct {
    model.Params
    // 入参pushProjectParam
    pushProjectParam   *ThirdProjectPushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushprojectRequest对象
func NewAlibabaDamaiMevOpenPushprojectRequest() *AlibabaDamaiMevOpenPushprojectRequest{
    return &AlibabaDamaiMevOpenPushprojectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushprojectRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushproject"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushprojectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushProjectParam Setter
// 入参pushProjectParam
func (r *AlibabaDamaiMevOpenPushprojectRequest) SetPushProjectParam(pushProjectParam *ThirdProjectPushOpenParam) error {
    r.pushProjectParam = pushProjectParam
    r.Set("push_project_param", pushProjectParam)
    return nil
}

// PushProjectParam Getter
func (r AlibabaDamaiMevOpenPushprojectRequest) GetPushProjectParam() *ThirdProjectPushOpenParam {
    return r.pushProjectParam
}
