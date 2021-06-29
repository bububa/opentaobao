package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后分期城市信息同步 APIRequest
tmall.car.lease.citysynchronize

天猫开新车租后分期城市信息同步
*/
type TmallCarLeaseCitysynchronizeRequest struct {
    model.Params

    // 地址信息
    paramAreaDto   *AreaDto 

}

func NewTmallCarLeaseCitysynchronizeRequest() *TmallCarLeaseCitysynchronizeRequest{
    return &TmallCarLeaseCitysynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseCitysynchronizeRequest) GetApiMethodName() string {
    return "tmall.car.lease.citysynchronize"
}

func (r TmallCarLeaseCitysynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseCitysynchronizeRequest) SetParamAreaDto(paramAreaDto *AreaDto) error {
    r.paramAreaDto = paramAreaDto
    r.Set("param_area_dto", paramAreaDto)
    return nil
}

func (r TmallCarLeaseCitysynchronizeRequest) GetParamAreaDto() *AreaDto {
    return r.paramAreaDto
}

