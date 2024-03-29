package ott

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttAlicbFacadeserviceGetdataAPIRequest 影视SDK获取设备能力值 API请求
// youku.ott.alicb.facadeservice.getdata
//
// 影视SDK获取设备能力值
type YoukuOttAlicbFacadeserviceGetdataAPIRequest struct {
	model.Params
	// 能力维度
	_serviceList []string
	// 设备唯一标识
	_uuid string
	// 属性MAP JSON串
	_propertyMapJson string
	// 扩展属性
	_extraInfoMap string
}

// NewYoukuOttAlicbFacadeserviceGetdataRequest 初始化YoukuOttAlicbFacadeserviceGetdataAPIRequest对象
func NewYoukuOttAlicbFacadeserviceGetdataRequest() *YoukuOttAlicbFacadeserviceGetdataAPIRequest {
	return &YoukuOttAlicbFacadeserviceGetdataAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttAlicbFacadeserviceGetdataAPIRequest) Reset() {
	r._serviceList = r._serviceList[:0]
	r._uuid = ""
	r._propertyMapJson = ""
	r._extraInfoMap = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttAlicbFacadeserviceGetdataAPIRequest) GetApiMethodName() string {
	return "youku.ott.alicb.facadeservice.getdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttAlicbFacadeserviceGetdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttAlicbFacadeserviceGetdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceList is ServiceList Setter
// 能力维度
func (r *YoukuOttAlicbFacadeserviceGetdataAPIRequest) SetServiceList(_serviceList []string) error {
	r._serviceList = _serviceList
	r.Set("service_list", _serviceList)
	return nil
}

// GetServiceList ServiceList Getter
func (r YoukuOttAlicbFacadeserviceGetdataAPIRequest) GetServiceList() []string {
	return r._serviceList
}

// SetUuid is Uuid Setter
// 设备唯一标识
func (r *YoukuOttAlicbFacadeserviceGetdataAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YoukuOttAlicbFacadeserviceGetdataAPIRequest) GetUuid() string {
	return r._uuid
}

// SetPropertyMapJson is PropertyMapJson Setter
// 属性MAP JSON串
func (r *YoukuOttAlicbFacadeserviceGetdataAPIRequest) SetPropertyMapJson(_propertyMapJson string) error {
	r._propertyMapJson = _propertyMapJson
	r.Set("property_map_json", _propertyMapJson)
	return nil
}

// GetPropertyMapJson PropertyMapJson Getter
func (r YoukuOttAlicbFacadeserviceGetdataAPIRequest) GetPropertyMapJson() string {
	return r._propertyMapJson
}

// SetExtraInfoMap is ExtraInfoMap Setter
// 扩展属性
func (r *YoukuOttAlicbFacadeserviceGetdataAPIRequest) SetExtraInfoMap(_extraInfoMap string) error {
	r._extraInfoMap = _extraInfoMap
	r.Set("extra_info_map", _extraInfoMap)
	return nil
}

// GetExtraInfoMap ExtraInfoMap Getter
func (r YoukuOttAlicbFacadeserviceGetdataAPIRequest) GetExtraInfoMap() string {
	return r._extraInfoMap
}

var poolYoukuOttAlicbFacadeserviceGetdataAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttAlicbFacadeserviceGetdataRequest()
	},
}

// GetYoukuOttAlicbFacadeserviceGetdataRequest 从 sync.Pool 获取 YoukuOttAlicbFacadeserviceGetdataAPIRequest
func GetYoukuOttAlicbFacadeserviceGetdataAPIRequest() *YoukuOttAlicbFacadeserviceGetdataAPIRequest {
	return poolYoukuOttAlicbFacadeserviceGetdataAPIRequest.Get().(*YoukuOttAlicbFacadeserviceGetdataAPIRequest)
}

// ReleaseYoukuOttAlicbFacadeserviceGetdataAPIRequest 将 YoukuOttAlicbFacadeserviceGetdataAPIRequest 放入 sync.Pool
func ReleaseYoukuOttAlicbFacadeserviceGetdataAPIRequest(v *YoukuOttAlicbFacadeserviceGetdataAPIRequest) {
	v.Reset()
	poolYoukuOttAlicbFacadeserviceGetdataAPIRequest.Put(v)
}
