package iotticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMaintainUpdateAPIRequest IoT售后服务商维修方案更新 API请求
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
type CainiaoIotTicketSpMaintainUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *UpdateMaintainPlanTopRequest
}

// NewCainiaoIotTicketSpMaintainUpdateRequest 初始化CainiaoIotTicketSpMaintainUpdateAPIRequest对象
func NewCainiaoIotTicketSpMaintainUpdateRequest() *CainiaoIotTicketSpMaintainUpdateAPIRequest {
	return &CainiaoIotTicketSpMaintainUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoIotTicketSpMaintainUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.maintain.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoIotTicketSpMaintainUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMaintainUpdateAPIRequest) SetParam(_param *UpdateMaintainPlanTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoIotTicketSpMaintainUpdateAPIRequest) GetParam() *UpdateMaintainPlanTopRequest {
	return r._param
}

var poolCainiaoIotTicketSpMaintainUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoIotTicketSpMaintainUpdateRequest()
	},
}

// GetCainiaoIotTicketSpMaintainUpdateRequest 从 sync.Pool 获取 CainiaoIotTicketSpMaintainUpdateAPIRequest
func GetCainiaoIotTicketSpMaintainUpdateAPIRequest() *CainiaoIotTicketSpMaintainUpdateAPIRequest {
	return poolCainiaoIotTicketSpMaintainUpdateAPIRequest.Get().(*CainiaoIotTicketSpMaintainUpdateAPIRequest)
}

// ReleaseCainiaoIotTicketSpMaintainUpdateAPIRequest 将 CainiaoIotTicketSpMaintainUpdateAPIRequest 放入 sync.Pool
func ReleaseCainiaoIotTicketSpMaintainUpdateAPIRequest(v *CainiaoIotTicketSpMaintainUpdateAPIRequest) {
	v.Reset()
	poolCainiaoIotTicketSpMaintainUpdateAPIRequest.Put(v)
}
