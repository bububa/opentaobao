package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest 根据uuid操作设备 API请求
// alibaba.campus.device.openapi.operatedevice
//
// 根据uuid操作设备
type AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest struct {
	model.Params
	// 请求发送端信息
	_workBenchContext *WorkBenchContext
	// 设备序列号uuid
	_uuid string
	// 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
	_propertyCode string
	// 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
	_value string
}

// NewAlibabaCampusDeviceOpenapiOperatedeviceRequest 初始化AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest对象
func NewAlibabaCampusDeviceOpenapiOperatedeviceRequest() *AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest {
	return &AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.operatedevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetUuid is Uuid Setter
// 设备序列号uuid
func (r *AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) GetUuid() string {
	return r._uuid
}

// SetPropertyCode is PropertyCode Setter
// 参数code,如灯亮度参数为brightness;设备的开关switchstate。参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’。
func (r *AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) SetPropertyCode(_propertyCode string) error {
	r._propertyCode = _propertyCode
	r.Set("property_code", _propertyCode)
	return nil
}

// GetPropertyCode PropertyCode Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) GetPropertyCode() string {
	return r._propertyCode
}

// SetValue is Value Setter
// 设置的参数值.如灯亮度为0~255.0表示关;设备开关,值使用on或off。[请按照‘设备详细信息开发文档’传入正确的参数值类型]
func (r *AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) SetValue(_value string) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r AlibabaCampusDeviceOpenapiOperatedeviceAPIRequest) GetValue() string {
	return r._value
}
