package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场馆接口deleteVenue APIRequest
alibaba.damai.mev.open.deletevenue

开放接口，删除场馆
*/
type AlibabaDamaiMevOpenDeletevenueRequest struct {
    model.Params

    // 入参deleteVenueParam
    deleteVenueParam   *VenueIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeletevenueRequest() *AlibabaDamaiMevOpenDeletevenueRequest{
    return &AlibabaDamaiMevOpenDeletevenueRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeletevenueRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deletevenue"
}

func (r AlibabaDamaiMevOpenDeletevenueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeletevenueRequest) SetDeleteVenueParam(deleteVenueParam *VenueIdOpenParam) error {
    r.deleteVenueParam = deleteVenueParam
    r.Set("delete_venue_param", deleteVenueParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeletevenueRequest) GetDeleteVenueParam() *VenueIdOpenParam {
    return r.deleteVenueParam
}

