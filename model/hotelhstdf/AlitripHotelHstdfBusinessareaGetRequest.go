package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据城市查询商圈 APIRequest
alitrip.hotel.hstdf.businessarea.get

根据cityId分页查询商圈信息
*/
type AlitripHotelHstdfBusinessareaGetRequest struct {
    model.Params

    // 请求参数封装
    paramGetByTrdiDivisionIdParam   *GetByTrdiDivisionIdParam 

}

func NewAlitripHotelHstdfBusinessareaGetRequest() *AlitripHotelHstdfBusinessareaGetRequest{
    return &AlitripHotelHstdfBusinessareaGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfBusinessareaGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.businessarea.get"
}

func (r AlitripHotelHstdfBusinessareaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfBusinessareaGetRequest) SetParamGetByTrdiDivisionIdParam(paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
    r.paramGetByTrdiDivisionIdParam = paramGetByTrdiDivisionIdParam
    r.Set("param_get_by_trdi_division_id_param", paramGetByTrdiDivisionIdParam)
    return nil
}

func (r AlitripHotelHstdfBusinessareaGetRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
    return r.paramGetByTrdiDivisionIdParam
}

