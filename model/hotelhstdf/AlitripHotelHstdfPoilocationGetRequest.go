package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据平台城市id分页查询poi location API请求
alitrip.hotel.hstdf.poilocation.get

根据平台城市id分页查询poi location
*/
type AlitripHotelHstdfPoilocationGetRequest struct {
    model.Params
    // 参数封装
    paramGetByTrdiDivisionIdParam   *GetByTrdiDivisionIdParam
}

// 初始化AlitripHotelHstdfPoilocationGetRequest对象
func NewAlitripHotelHstdfPoilocationGetRequest() *AlitripHotelHstdfPoilocationGetRequest{
    return &AlitripHotelHstdfPoilocationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfPoilocationGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.poilocation.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfPoilocationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamGetByTrdiDivisionIdParam Setter
// 参数封装
func (r *AlitripHotelHstdfPoilocationGetRequest) SetParamGetByTrdiDivisionIdParam(paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
    r.paramGetByTrdiDivisionIdParam = paramGetByTrdiDivisionIdParam
    r.Set("param_get_by_trdi_division_id_param", paramGetByTrdiDivisionIdParam)
    return nil
}

// ParamGetByTrdiDivisionIdParam Getter
func (r AlitripHotelHstdfPoilocationGetRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
    return r.paramGetByTrdiDivisionIdParam
}
