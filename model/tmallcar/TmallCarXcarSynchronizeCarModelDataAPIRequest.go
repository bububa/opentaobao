package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarXcarSynchronizeCarModelDataAPIRequest 爱车车型数据同步 API请求
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
type TmallCarXcarSynchronizeCarModelDataAPIRequest struct {
	model.Params
	// 传入对象描述
	_paramXCarSysModelDTO *XcarSysModelDto
}

// NewTmallCarXcarSynchronizeCarModelDataRequest 初始化TmallCarXcarSynchronizeCarModelDataAPIRequest对象
func NewTmallCarXcarSynchronizeCarModelDataRequest() *TmallCarXcarSynchronizeCarModelDataAPIRequest {
	return &TmallCarXcarSynchronizeCarModelDataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarXcarSynchronizeCarModelDataAPIRequest) GetApiMethodName() string {
	return "tmall.car.xcar.synchronize.car.model.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarXcarSynchronizeCarModelDataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarXcarSynchronizeCarModelDataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamXCarSysModelDTO is ParamXCarSysModelDTO Setter
// 传入对象描述
func (r *TmallCarXcarSynchronizeCarModelDataAPIRequest) SetParamXCarSysModelDTO(_paramXCarSysModelDTO *XcarSysModelDto) error {
	r._paramXCarSysModelDTO = _paramXCarSysModelDTO
	r.Set("param_x_car_sys_model_d_t_o", _paramXCarSysModelDTO)
	return nil
}

// GetParamXCarSysModelDTO ParamXCarSysModelDTO Getter
func (r TmallCarXcarSynchronizeCarModelDataAPIRequest) GetParamXCarSysModelDTO() *XcarSysModelDto {
	return r._paramXCarSysModelDTO
}
