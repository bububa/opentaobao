package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据平台城市id分页查询poi location APIRequest
alitrip.hotel.hstdf.poilocation.get

根据平台城市id分页查询poi location
*/
type AlitripHotelHstdfPoilocationGetRequest struct {
    model.Params

    // 参数封装
    paramGetByTrdiDivisionIdParam   *GetByTrdiDivisionIdParam 

}

func NewAlitripHotelHstdfPoilocationGetRequest() *AlitripHotelHstdfPoilocationGetRequest{
    return &AlitripHotelHstdfPoilocationGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfPoilocationGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.poilocation.get"
}

func (r AlitripHotelHstdfPoilocationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfPoilocationGetRequest) SetParamGetByTrdiDivisionIdParam(paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
    r.paramGetByTrdiDivisionIdParam = paramGetByTrdiDivisionIdParam
    r.Set("param_get_by_trdi_division_id_param", paramGetByTrdiDivisionIdParam)
    return nil
}

func (r AlitripHotelHstdfPoilocationGetRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
    return r.paramGetByTrdiDivisionIdParam
}

