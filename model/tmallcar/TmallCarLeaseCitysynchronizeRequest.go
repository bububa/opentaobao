package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后分期城市信息同步 API请求
tmall.car.lease.citysynchronize

天猫开新车租后分期城市信息同步
*/
type TmallCarLeaseCitysynchronizeRequest struct {
    model.Params
    // 地址信息
    paramAreaDto   *AreaDto
}

// 初始化TmallCarLeaseCitysynchronizeRequest对象
func NewTmallCarLeaseCitysynchronizeRequest() *TmallCarLeaseCitysynchronizeRequest{
    return &TmallCarLeaseCitysynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseCitysynchronizeRequest) GetApiMethodName() string {
    return "tmall.car.lease.citysynchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseCitysynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAreaDto Setter
// 地址信息
func (r *TmallCarLeaseCitysynchronizeRequest) SetParamAreaDto(paramAreaDto *AreaDto) error {
    r.paramAreaDto = paramAreaDto
    r.Set("param_area_dto", paramAreaDto)
    return nil
}

// ParamAreaDto Getter
func (r TmallCarLeaseCitysynchronizeRequest) GetParamAreaDto() *AreaDto {
    return r.paramAreaDto
}
