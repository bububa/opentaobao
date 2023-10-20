package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest saveeventinfoforibos API请求
// alibaba.campus.device.openapi.saveeventinfoforibos
//
// IBos的事件信息上报与反馈的处理接口
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 系统自动生成
	_param1 *EventInfoApiDto
}

// NewAlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest 初始化AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest对象
func NewAlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest() *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest {
	return &AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.saveeventinfoforibos"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) SetParam1(_param1 *EventInfoApiDto) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetParam1() *EventInfoApiDto {
	return r._param1
}

var poolAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest
func GetAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest() *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest {
	return poolAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest.Get().(*AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest 将 AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest(v *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest.Put(v)
}
