package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleasecitysynchronizeAPIRequest 天猫开新车租后分期城市信息同步 API请求
// tmall.car.lease.citysynchronize
//
// 天猫开新车租后分期城市信息同步
type TmallcarleasecitysynchronizeAPIRequest struct {
	model.Params
	// 地址信息
	_paramAreaDto *AreaDto
}

// NewTmallcarleasecitysynchronizeRequest 初始化TmallcarleasecitysynchronizeAPIRequest对象
func NewTmallcarleasecitysynchronizeRequest() *TmallcarleasecitysynchronizeAPIRequest {
	return &TmallcarleasecitysynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleasecitysynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.citysynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleasecitysynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleasecitysynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAreaDto is ParamAreaDto Setter
// 地址信息
func (r *TmallcarleasecitysynchronizeAPIRequest) SetParamAreaDto(_paramAreaDto *AreaDto) error {
	r._paramAreaDto = _paramAreaDto
	r.Set("param_area_dto", _paramAreaDto)
	return nil
}

// GetParamAreaDto ParamAreaDto Getter
func (r TmallcarleasecitysynchronizeAPIRequest) GetParamAreaDto() *AreaDto {
	return r._paramAreaDto
}
