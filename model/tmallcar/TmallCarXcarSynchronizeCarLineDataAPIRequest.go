package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarXcarSynchronizeCarLineDataAPIRequest 我的爱卡车型配置数据 API请求
// tmall.car.xcar.synchronize.car.line.data
//
// 同步我的爱卡车系配置数据到天猫汽车
type TmallCarXcarSynchronizeCarLineDataAPIRequest struct {
	model.Params
	// 入参对象
	_paramXCarSysLineDTO *XCarSysLineDto
}

// NewTmallCarXcarSynchronizeCarLineDataRequest 初始化TmallCarXcarSynchronizeCarLineDataAPIRequest对象
func NewTmallCarXcarSynchronizeCarLineDataRequest() *TmallCarXcarSynchronizeCarLineDataAPIRequest {
	return &TmallCarXcarSynchronizeCarLineDataAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarXcarSynchronizeCarLineDataAPIRequest) Reset() {
	r._paramXCarSysLineDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarXcarSynchronizeCarLineDataAPIRequest) GetApiMethodName() string {
	return "tmall.car.xcar.synchronize.car.line.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarXcarSynchronizeCarLineDataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarXcarSynchronizeCarLineDataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamXCarSysLineDTO is ParamXCarSysLineDTO Setter
// 入参对象
func (r *TmallCarXcarSynchronizeCarLineDataAPIRequest) SetParamXCarSysLineDTO(_paramXCarSysLineDTO *XCarSysLineDto) error {
	r._paramXCarSysLineDTO = _paramXCarSysLineDTO
	r.Set("param_x_car_sys_line_d_t_o", _paramXCarSysLineDTO)
	return nil
}

// GetParamXCarSysLineDTO ParamXCarSysLineDTO Getter
func (r TmallCarXcarSynchronizeCarLineDataAPIRequest) GetParamXCarSysLineDTO() *XCarSysLineDto {
	return r._paramXCarSysLineDTO
}

var poolTmallCarXcarSynchronizeCarLineDataAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarXcarSynchronizeCarLineDataRequest()
	},
}

// GetTmallCarXcarSynchronizeCarLineDataRequest 从 sync.Pool 获取 TmallCarXcarSynchronizeCarLineDataAPIRequest
func GetTmallCarXcarSynchronizeCarLineDataAPIRequest() *TmallCarXcarSynchronizeCarLineDataAPIRequest {
	return poolTmallCarXcarSynchronizeCarLineDataAPIRequest.Get().(*TmallCarXcarSynchronizeCarLineDataAPIRequest)
}

// ReleaseTmallCarXcarSynchronizeCarLineDataAPIRequest 将 TmallCarXcarSynchronizeCarLineDataAPIRequest 放入 sync.Pool
func ReleaseTmallCarXcarSynchronizeCarLineDataAPIRequest(v *TmallCarXcarSynchronizeCarLineDataAPIRequest) {
	v.Reset()
	poolTmallCarXcarSynchronizeCarLineDataAPIRequest.Put(v)
}
