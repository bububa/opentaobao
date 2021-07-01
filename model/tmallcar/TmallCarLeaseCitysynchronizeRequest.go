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
type TmallCarLeaseCitysynchronizeAPIRequest struct {
    model.Params
    // 地址信息
    _paramAreaDto   *AreaDTO
}

// 初始化TmallCarLeaseCitysynchronizeAPIRequest对象
func NewTmallCarLeaseCitysynchronizeRequest() *TmallCarLeaseCitysynchronizeAPIRequest{
    return &TmallCarLeaseCitysynchronizeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseCitysynchronizeAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.citysynchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseCitysynchronizeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAreaDto Setter
// 地址信息
func (r *TmallCarLeaseCitysynchronizeAPIRequest) SetParamAreaDto(_paramAreaDto *AreaDTO) error {
    r._paramAreaDto = _paramAreaDto
    r.Set("param_area_dto", _paramAreaDto)
    return nil
}

// ParamAreaDto Getter
func (r TmallCarLeaseCitysynchronizeAPIRequest) GetParamAreaDto() *AreaDTO {
    return r._paramAreaDto
}
