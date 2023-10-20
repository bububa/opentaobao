package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosPosAlarmAPIRequest pos故障报警 API请求
// alibaba.mos.pos.alarm
//
// 故障报警
type AlibabaMosPosAlarmAPIRequest struct {
	model.Params
	// 参数
	_param0 *PosLogDto
}

// NewAlibabaMosPosAlarmRequest 初始化AlibabaMosPosAlarmAPIRequest对象
func NewAlibabaMosPosAlarmRequest() *AlibabaMosPosAlarmAPIRequest {
	return &AlibabaMosPosAlarmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosPosAlarmAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosPosAlarmAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.pos.alarm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosPosAlarmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosPosAlarmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabaMosPosAlarmAPIRequest) SetParam0(_param0 *PosLogDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaMosPosAlarmAPIRequest) GetParam0() *PosLogDto {
	return r._param0
}

var poolAlibabaMosPosAlarmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosPosAlarmRequest()
	},
}

// GetAlibabaMosPosAlarmRequest 从 sync.Pool 获取 AlibabaMosPosAlarmAPIRequest
func GetAlibabaMosPosAlarmAPIRequest() *AlibabaMosPosAlarmAPIRequest {
	return poolAlibabaMosPosAlarmAPIRequest.Get().(*AlibabaMosPosAlarmAPIRequest)
}

// ReleaseAlibabaMosPosAlarmAPIRequest 将 AlibabaMosPosAlarmAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosPosAlarmAPIRequest(v *AlibabaMosPosAlarmAPIRequest) {
	v.Reset()
	poolAlibabaMosPosAlarmAPIRequest.Put(v)
}
