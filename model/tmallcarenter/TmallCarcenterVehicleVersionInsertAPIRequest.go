package tmallcarenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarcenterVehicleVersionInsertAPIRequest
汽车EPC版本压缩库新增接口 API请求
tmall.carcenter.vehicle.version.insert

汽车EPC版本压缩库新增接口 */
type TmallCarcenterVehicleVersionInsertAPIRequest struct {
	model.Params
	// 版本压缩库入参
	_dto *VersionVehicleInfoOriginalDto
}

// NewTmallCarcenterVehicleVersionInsertRequest 初始化TmallCarcenterVehicleVersionInsertAPIRequest对象
func NewTmallCarcenterVehicleVersionInsertRequest() *TmallCarcenterVehicleVersionInsertAPIRequest {
	return &TmallCarcenterVehicleVersionInsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleVersionInsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.version.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleVersionInsertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Dto Setter
// 版本压缩库入参
func (r *TmallCarcenterVehicleVersionInsertAPIRequest) SetDto(_dto *VersionVehicleInfoOriginalDto) error {
	r._dto = _dto
	r.Set("dto", _dto)
	return nil
}

// Get Dto Getter
func (r TmallCarcenterVehicleVersionInsertAPIRequest) GetDto() *VersionVehicleInfoOriginalDto {
	return r._dto
}
