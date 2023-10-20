package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapioperatedeviceAPIRequest 根据uuid操作设备 API请求
// alibaba.campus.device.openapi.operatedevice
//
// 根据uuid操作设备
type AlibabacampusdeviceopenapioperatedeviceAPIRequest struct {
	model.Params
	// 设备序列号uuid
	_uuid string
	// 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
	_propertyCode string
	// 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
	_value string
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabacampusdeviceopenapioperatedeviceRequest 初始化AlibabacampusdeviceopenapioperatedeviceAPIRequest对象
func NewAlibabacampusdeviceopenapioperatedeviceRequest() *AlibabacampusdeviceopenapioperatedeviceAPIRequest {
	return &AlibabacampusdeviceopenapioperatedeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapioperatedeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.operatedevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapioperatedeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapioperatedeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备序列号uuid
func (r *AlibabacampusdeviceopenapioperatedeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabacampusdeviceopenapioperatedeviceAPIRequest) GetUuid() string {
	return r._uuid
}

// SetPropertyCode is PropertyCode Setter
// 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
func (r *AlibabacampusdeviceopenapioperatedeviceAPIRequest) SetPropertyCode(_propertyCode string) error {
	r._propertyCode = _propertyCode
	r.Set("property_code", _propertyCode)
	return nil
}

// GetPropertyCode PropertyCode Getter
func (r AlibabacampusdeviceopenapioperatedeviceAPIRequest) GetPropertyCode() string {
	return r._propertyCode
}

// SetValue is Value Setter
// 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
func (r *AlibabacampusdeviceopenapioperatedeviceAPIRequest) SetValue(_value string) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r AlibabacampusdeviceopenapioperatedeviceAPIRequest) GetValue() string {
	return r._value
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabacampusdeviceopenapioperatedeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapioperatedeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}
