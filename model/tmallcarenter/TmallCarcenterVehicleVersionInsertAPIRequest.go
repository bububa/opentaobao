package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcentervehicleversioninsertAPIRequest 汽车EPC版本压缩库新增接口 API请求
// tmall.carcenter.vehicle.version.insert
//
// 汽车EPC版本压缩库新增接口
type TmallcarcentervehicleversioninsertAPIRequest struct {
	model.Params
	// 版本压缩库入参
	_dto *VersionVehicleInfoOriginalDto
}

// NewTmallcarcentervehicleversioninsertRequest 初始化TmallcarcentervehicleversioninsertAPIRequest对象
func NewTmallcarcentervehicleversioninsertRequest() *TmallcarcentervehicleversioninsertAPIRequest {
	return &TmallcarcentervehicleversioninsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarcentervehicleversioninsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.version.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarcentervehicleversioninsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarcentervehicleversioninsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDto is Dto Setter
// 版本压缩库入参
func (r *TmallcarcentervehicleversioninsertAPIRequest) SetDto(_dto *VersionVehicleInfoOriginalDto) error {
	r._dto = _dto
	r.Set("dto", _dto)
	return nil
}

// GetDto Dto Getter
func (r TmallcarcentervehicleversioninsertAPIRequest) GetDto() *VersionVehicleInfoOriginalDto {
	return r._dto
}
