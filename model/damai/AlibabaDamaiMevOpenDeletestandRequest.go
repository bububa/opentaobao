package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-看台接口deleteStand API请求
alibaba.damai.mev.open.deletestand

deleteStand
*/
type AlibabaDamaiMevOpenDeletestandRequest struct {
    model.Params
    // 入参deleteStandParam
    _deleteStandParam   *StandIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeletestandRequest对象
func NewAlibabaDamaiMevOpenDeletestandRequest() *AlibabaDamaiMevOpenDeletestandRequest{
    return &AlibabaDamaiMevOpenDeletestandRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletestandRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deletestand"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletestandRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteStandParam Setter
// 入参deleteStandParam
func (r *AlibabaDamaiMevOpenDeletestandRequest) SetDeleteStandParam(_deleteStandParam *StandIdOpenParam) error {
    r._deleteStandParam = _deleteStandParam
    r.Set("delete_stand_param", _deleteStandParam)
    return nil
}

// DeleteStandParam Getter
func (r AlibabaDamaiMevOpenDeletestandRequest) GetDeleteStandParam() *StandIdOpenParam {
    return r._deleteStandParam
}
