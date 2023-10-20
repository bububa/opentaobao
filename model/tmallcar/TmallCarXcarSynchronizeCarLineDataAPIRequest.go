package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarxcarsynchronizecarlinedataAPIRequest 我的爱卡车型配置数据 API请求
// tmall.car.xcar.synchronize.car.line.data
//
// 同步我的爱卡车系配置数据到天猫汽车
type TmallcarxcarsynchronizecarlinedataAPIRequest struct {
	model.Params
	// 入参对象
	_paramXCarSysLineDTO *XcarSysLineDto
}

// NewTmallcarxcarsynchronizecarlinedataRequest 初始化TmallcarxcarsynchronizecarlinedataAPIRequest对象
func NewTmallcarxcarsynchronizecarlinedataRequest() *TmallcarxcarsynchronizecarlinedataAPIRequest {
	return &TmallcarxcarsynchronizecarlinedataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarxcarsynchronizecarlinedataAPIRequest) GetApiMethodName() string {
	return "tmall.car.xcar.synchronize.car.line.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarxcarsynchronizecarlinedataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarxcarsynchronizecarlinedataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamXCarSysLineDTO is ParamXCarSysLineDTO Setter
// 入参对象
func (r *TmallcarxcarsynchronizecarlinedataAPIRequest) SetParamXCarSysLineDTO(_paramXCarSysLineDTO *XcarSysLineDto) error {
	r._paramXCarSysLineDTO = _paramXCarSysLineDTO
	r.Set("param_x_car_sys_line_d_t_o", _paramXCarSysLineDTO)
	return nil
}

// GetParamXCarSysLineDTO ParamXCarSysLineDTO Getter
func (r TmallcarxcarsynchronizecarlinedataAPIRequest) GetParamXCarSysLineDTO() *XcarSysLineDto {
	return r._paramXCarSysLineDTO
}
