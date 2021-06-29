package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-看台接口deleteStand APIRequest
alibaba.damai.mev.open.deletestand

deleteStand
*/
type AlibabaDamaiMevOpenDeletestandRequest struct {
    model.Params

    // 入参deleteStandParam
    deleteStandParam   *StandIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeletestandRequest() *AlibabaDamaiMevOpenDeletestandRequest{
    return &AlibabaDamaiMevOpenDeletestandRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeletestandRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deletestand"
}

func (r AlibabaDamaiMevOpenDeletestandRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeletestandRequest) SetDeleteStandParam(deleteStandParam *StandIdOpenParam) error {
    r.deleteStandParam = deleteStandParam
    r.Set("delete_stand_param", deleteStandParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeletestandRequest) GetDeleteStandParam() *StandIdOpenParam {
    return r.deleteStandParam
}

