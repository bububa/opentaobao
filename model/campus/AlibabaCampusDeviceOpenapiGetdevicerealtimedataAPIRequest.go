package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest 获取指定设备下指定参数的实时值 API请求
// alibaba.campus.device.openapi.getdevicerealtimedata
//
// 获取指定设备下指定参数的实时值
type AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest struct {
	model.Params
	// 设备uuid
	_uuid string
	// 参数code,如灯亮度参数为brightness;参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’[根据设备类型查看该设备所拥有的采集类参数]。
	_propertyCode string
	// 请求端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest 初始化AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest {
	return &AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) Reset() {
	r._uuid = ""
	r._propertyCode = ""
	r._workBenchContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getdevicerealtimedata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) GetUuid() string {
	return r._uuid
}

// SetPropertyCode is PropertyCode Setter
// 参数code,如灯亮度参数为brightness;参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’[根据设备类型查看该设备所拥有的采集类参数]。
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) SetPropertyCode(_propertyCode string) error {
	r._propertyCode = _propertyCode
	r.Set("property_code", _propertyCode)
	return nil
}

// GetPropertyCode PropertyCode Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) GetPropertyCode() string {
	return r._propertyCode
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求端信息
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

var poolAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest()
	},
}

// GetAlibabaCampusDeviceOpenapiGetdevicerealtimedataRequest 从 sync.Pool 获取 AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest
func GetAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest {
	return poolAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest.Get().(*AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest)
}

// ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest 将 AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest(v *AlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest) {
	v.Reset()
	poolAlibabaCampusDeviceOpenapiGetdevicerealtimedataAPIRequest.Put(v)
}
