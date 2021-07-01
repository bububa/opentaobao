package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场次接口deletePerform API请求
alibaba.damai.mev.open.deleteperform

deletePerform
*/
type AlibabaDamaiMevOpenDeleteperformAPIRequest struct {
    model.Params
    // 入参deletePerformParam
    _deletePerformParam   *PerformIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeleteperformAPIRequest对象
func NewAlibabaDamaiMevOpenDeleteperformRequest() *AlibabaDamaiMevOpenDeleteperformAPIRequest{
    return &AlibabaDamaiMevOpenDeleteperformAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteperformAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteperform"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteperformAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeletePerformParam Setter
// 入参deletePerformParam
func (r *AlibabaDamaiMevOpenDeleteperformAPIRequest) SetDeletePerformParam(_deletePerformParam *PerformIdOpenParam) error {
    r._deletePerformParam = _deletePerformParam
    r.Set("delete_perform_param", _deletePerformParam)
    return nil
}

// DeletePerformParam Getter
func (r AlibabaDamaiMevOpenDeleteperformAPIRequest) GetDeletePerformParam() *PerformIdOpenParam {
    return r._deletePerformParam
}
