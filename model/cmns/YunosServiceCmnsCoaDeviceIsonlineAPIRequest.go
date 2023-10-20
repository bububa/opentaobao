package cmns

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaDeviceIsonlineAPIRequest 根据设备id查询设备是否在线 API请求
// yunos.service.cmns.coa.device.isonline
//
// 根据设备id查询设备是否在线
type YunosServiceCmnsCoaDeviceIsonlineAPIRequest struct {
	model.Params
	// 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
	_type string
	// 对应的设备id值
	_value string
}

// NewYunosServiceCmnsCoaDeviceIsonlineRequest 初始化YunosServiceCmnsCoaDeviceIsonlineAPIRequest对象
func NewYunosServiceCmnsCoaDeviceIsonlineRequest() *YunosServiceCmnsCoaDeviceIsonlineAPIRequest {
	return &YunosServiceCmnsCoaDeviceIsonlineAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosServiceCmnsCoaDeviceIsonlineAPIRequest) Reset() {
	r._type = ""
	r._value = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.device.isonline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 设备id类型，取值&#34;uuid&#34;或者&#34;imei&#34;或者&#34;deviceToken&#34;
func (r *YunosServiceCmnsCoaDeviceIsonlineAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetType() string {
	return r._type
}

// SetValue is Value Setter
// 对应的设备id值
func (r *YunosServiceCmnsCoaDeviceIsonlineAPIRequest) SetValue(_value string) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r YunosServiceCmnsCoaDeviceIsonlineAPIRequest) GetValue() string {
	return r._value
}

var poolYunosServiceCmnsCoaDeviceIsonlineAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosServiceCmnsCoaDeviceIsonlineRequest()
	},
}

// GetYunosServiceCmnsCoaDeviceIsonlineRequest 从 sync.Pool 获取 YunosServiceCmnsCoaDeviceIsonlineAPIRequest
func GetYunosServiceCmnsCoaDeviceIsonlineAPIRequest() *YunosServiceCmnsCoaDeviceIsonlineAPIRequest {
	return poolYunosServiceCmnsCoaDeviceIsonlineAPIRequest.Get().(*YunosServiceCmnsCoaDeviceIsonlineAPIRequest)
}

// ReleaseYunosServiceCmnsCoaDeviceIsonlineAPIRequest 将 YunosServiceCmnsCoaDeviceIsonlineAPIRequest 放入 sync.Pool
func ReleaseYunosServiceCmnsCoaDeviceIsonlineAPIRequest(v *YunosServiceCmnsCoaDeviceIsonlineAPIRequest) {
	v.Reset()
	poolYunosServiceCmnsCoaDeviceIsonlineAPIRequest.Put(v)
}
