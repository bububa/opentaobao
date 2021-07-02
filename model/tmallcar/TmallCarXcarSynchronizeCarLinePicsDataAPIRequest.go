package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarXcarSynchronizeCarLinePicsDataAPIRequest 爱卡车系图片数据接入 API请求
// tmall.car.xcar.synchronize.car.line.pics.data
//
// 爱卡车系图片数据同步天猫汽车
type TmallCarXcarSynchronizeCarLinePicsDataAPIRequest struct {
	model.Params
	// 入参对象
	_paramXCarSysLinePicsDTO *XCarSysLinePicsDto
}

// NewTmallCarXcarSynchronizeCarLinePicsDataRequest 初始化TmallCarXcarSynchronizeCarLinePicsDataAPIRequest对象
func NewTmallCarXcarSynchronizeCarLinePicsDataRequest() *TmallCarXcarSynchronizeCarLinePicsDataAPIRequest {
	return &TmallCarXcarSynchronizeCarLinePicsDataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarXcarSynchronizeCarLinePicsDataAPIRequest) GetApiMethodName() string {
	return "tmall.car.xcar.synchronize.car.line.pics.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarXcarSynchronizeCarLinePicsDataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamXCarSysLinePicsDTO is ParamXCarSysLinePicsDTO Setter
// 入参对象
func (r *TmallCarXcarSynchronizeCarLinePicsDataAPIRequest) SetParamXCarSysLinePicsDTO(_paramXCarSysLinePicsDTO *XCarSysLinePicsDto) error {
	r._paramXCarSysLinePicsDTO = _paramXCarSysLinePicsDTO
	r.Set("param_x_car_sys_line_pics_d_t_o", _paramXCarSysLinePicsDTO)
	return nil
}

// GetParamXCarSysLinePicsDTO ParamXCarSysLinePicsDTO Getter
func (r TmallCarXcarSynchronizeCarLinePicsDataAPIRequest) GetParamXCarSysLinePicsDTO() *XCarSysLinePicsDto {
	return r._paramXCarSysLinePicsDTO
}
