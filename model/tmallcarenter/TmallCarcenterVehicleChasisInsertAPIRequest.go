package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleChasisInsertAPIRequest EPC车型底盘压缩库新增接口 API请求
// tmall.carcenter.vehicle.chasis.insert
//
// EPC车型底盘压缩库新增接口
type TmallCarcenterVehicleChasisInsertAPIRequest struct {
	model.Params
	// 底盘压缩库入参
	_dto *ChasisVehicleInfoOriginalDto
}

// NewTmallCarcenterVehicleChasisInsertRequest 初始化TmallCarcenterVehicleChasisInsertAPIRequest对象
func NewTmallCarcenterVehicleChasisInsertRequest() *TmallCarcenterVehicleChasisInsertAPIRequest {
	return &TmallCarcenterVehicleChasisInsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleChasisInsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.chasis.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleChasisInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarcenterVehicleChasisInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDto is Dto Setter
// 底盘压缩库入参
func (r *TmallCarcenterVehicleChasisInsertAPIRequest) SetDto(_dto *ChasisVehicleInfoOriginalDto) error {
	r._dto = _dto
	r.Set("dto", _dto)
	return nil
}

// GetDto Dto Getter
func (r TmallCarcenterVehicleChasisInsertAPIRequest) GetDto() *ChasisVehicleInfoOriginalDto {
	return r._dto
}
