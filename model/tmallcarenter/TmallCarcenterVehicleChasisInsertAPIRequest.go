package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcentervehiclechasisinsertAPIRequest EPC车型底盘压缩库新增接口 API请求
// tmall.carcenter.vehicle.chasis.insert
//
// EPC车型底盘压缩库新增接口
type TmallcarcentervehiclechasisinsertAPIRequest struct {
	model.Params
	// 底盘压缩库入参
	_dto *ChasisVehicleInfoOriginalDto
}

// NewTmallcarcentervehiclechasisinsertRequest 初始化TmallcarcentervehiclechasisinsertAPIRequest对象
func NewTmallcarcentervehiclechasisinsertRequest() *TmallcarcentervehiclechasisinsertAPIRequest {
	return &TmallcarcentervehiclechasisinsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarcentervehiclechasisinsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.chasis.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarcentervehiclechasisinsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarcentervehiclechasisinsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDto is Dto Setter
// 底盘压缩库入参
func (r *TmallcarcentervehiclechasisinsertAPIRequest) SetDto(_dto *ChasisVehicleInfoOriginalDto) error {
	r._dto = _dto
	r.Set("dto", _dto)
	return nil
}

// GetDto Dto Getter
func (r TmallcarcentervehiclechasisinsertAPIRequest) GetDto() *ChasisVehicleInfoOriginalDto {
	return r._dto
}
