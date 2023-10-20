package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest 获取指定设备下指定参数的实时值 API请求
// alibaba.campus.device.openapi.getdevicerealtimedata
//
// 获取指定设备下指定参数的实时值
type AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest struct {
	model.Params
	// 设备uuid
	_uuid string
	// 参数code,如灯亮度参数为brightness;参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’[根据设备类型查看该设备所拥有的采集类参数]。
	_propertyCode string
	// 请求端信息
	_workBenchContext *WorkBenchContext
}

// NewAlibabacampusdeviceopenapigetdevicerealtimedataRequest 初始化AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest对象
func NewAlibabacampusdeviceopenapigetdevicerealtimedataRequest() *AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest {
	return &AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.device.openapi.getdevicerealtimedata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) GetUuid() string {
	return r._uuid
}

// SetPropertyCode is PropertyCode Setter
// 参数code,如灯亮度参数为brightness;参数code信息请查阅‘平台技术’下‘设备详细信息开发文档’[根据设备类型查看该设备所拥有的采集类参数]。
func (r *AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) SetPropertyCode(_propertyCode string) error {
	r._propertyCode = _propertyCode
	r.Set("property_code", _propertyCode)
	return nil
}

// GetPropertyCode PropertyCode Getter
func (r AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) GetPropertyCode() string {
	return r._propertyCode
}

// SetWorkBenchContext is WorkBenchContext Setter
// 请求端信息
func (r *AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}
