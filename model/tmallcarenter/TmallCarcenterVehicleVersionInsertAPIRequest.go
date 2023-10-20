package tmallcarenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleVersionInsertAPIRequest 汽车EPC版本压缩库新增接口 API请求
// tmall.carcenter.vehicle.version.insert
//
// 汽车EPC版本压缩库新增接口
type TmallCarcenterVehicleVersionInsertAPIRequest struct {
	model.Params
	// 版本压缩库入参
	_dto *VersionVehicleInfoOriginalDto
}

// NewTmallCarcenterVehicleVersionInsertRequest 初始化TmallCarcenterVehicleVersionInsertAPIRequest对象
func NewTmallCarcenterVehicleVersionInsertRequest() *TmallCarcenterVehicleVersionInsertAPIRequest {
	return &TmallCarcenterVehicleVersionInsertAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarcenterVehicleVersionInsertAPIRequest) Reset() {
	r._dto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleVersionInsertAPIRequest) GetApiMethodName() string {
	return "tmall.carcenter.vehicle.version.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleVersionInsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarcenterVehicleVersionInsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDto is Dto Setter
// 版本压缩库入参
func (r *TmallCarcenterVehicleVersionInsertAPIRequest) SetDto(_dto *VersionVehicleInfoOriginalDto) error {
	r._dto = _dto
	r.Set("dto", _dto)
	return nil
}

// GetDto Dto Getter
func (r TmallCarcenterVehicleVersionInsertAPIRequest) GetDto() *VersionVehicleInfoOriginalDto {
	return r._dto
}

var poolTmallCarcenterVehicleVersionInsertAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarcenterVehicleVersionInsertRequest()
	},
}

// GetTmallCarcenterVehicleVersionInsertRequest 从 sync.Pool 获取 TmallCarcenterVehicleVersionInsertAPIRequest
func GetTmallCarcenterVehicleVersionInsertAPIRequest() *TmallCarcenterVehicleVersionInsertAPIRequest {
	return poolTmallCarcenterVehicleVersionInsertAPIRequest.Get().(*TmallCarcenterVehicleVersionInsertAPIRequest)
}

// ReleaseTmallCarcenterVehicleVersionInsertAPIRequest 将 TmallCarcenterVehicleVersionInsertAPIRequest 放入 sync.Pool
func ReleaseTmallCarcenterVehicleVersionInsertAPIRequest(v *TmallCarcenterVehicleVersionInsertAPIRequest) {
	v.Reset()
	poolTmallCarcenterVehicleVersionInsertAPIRequest.Put(v)
}
