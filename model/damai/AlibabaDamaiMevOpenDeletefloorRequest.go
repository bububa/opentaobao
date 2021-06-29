package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-楼层接口deleteFloor APIRequest
alibaba.damai.mev.open.deletefloor

deleteFloor
*/
type AlibabaDamaiMevOpenDeletefloorRequest struct {
    model.Params

    // 入参deleteFloorParam
    deleteFloorParam   *FloorIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeletefloorRequest() *AlibabaDamaiMevOpenDeletefloorRequest{
    return &AlibabaDamaiMevOpenDeletefloorRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeletefloorRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deletefloor"
}

func (r AlibabaDamaiMevOpenDeletefloorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeletefloorRequest) SetDeleteFloorParam(deleteFloorParam *FloorIdOpenParam) error {
    r.deleteFloorParam = deleteFloorParam
    r.Set("delete_floor_param", deleteFloorParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeletefloorRequest) GetDeleteFloorParam() *FloorIdOpenParam {
    return r.deleteFloorParam
}

