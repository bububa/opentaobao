package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据城市查询商圈 API请求
alitrip.hotel.hstdf.businessarea.get

根据cityId分页查询商圈信息
*/
type AlitripHotelHstdfBusinessareaGetAPIRequest struct {
    model.Params
    // 请求参数封装
    _paramGetByTrdiDivisionIdParam   *GetByTrdiDivisionIdParam
}

// 初始化AlitripHotelHstdfBusinessareaGetAPIRequest对象
func NewAlitripHotelHstdfBusinessareaGetRequest() *AlitripHotelHstdfBusinessareaGetAPIRequest{
    return &AlitripHotelHstdfBusinessareaGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfBusinessareaGetAPIRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.businessarea.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfBusinessareaGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamGetByTrdiDivisionIdParam Setter
// 请求参数封装
func (r *AlitripHotelHstdfBusinessareaGetAPIRequest) SetParamGetByTrdiDivisionIdParam(_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
    r._paramGetByTrdiDivisionIdParam = _paramGetByTrdiDivisionIdParam
    r.Set("param_get_by_trdi_division_id_param", _paramGetByTrdiDivisionIdParam)
    return nil
}

// ParamGetByTrdiDivisionIdParam Getter
func (r AlitripHotelHstdfBusinessareaGetAPIRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
    return r._paramGetByTrdiDivisionIdParam
}
