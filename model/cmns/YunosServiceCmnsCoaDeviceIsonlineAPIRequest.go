package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoadeviceisonlineAPIRequest 根据设备id查询设备是否在线 API请求
// yunos.service.cmns.coa.device.isonline
//
// 根据设备id查询设备是否在线
type YunosservicecmnscoadeviceisonlineAPIRequest struct {
	model.Params
	// 设备id类型，取值"uuid"或者"imei"或者"deviceToken"
	_type string
	// 对应的设备id值
	_value string
}

// NewYunosservicecmnscoadeviceisonlineRequest 初始化YunosservicecmnscoadeviceisonlineAPIRequest对象
func NewYunosservicecmnscoadeviceisonlineRequest() *YunosservicecmnscoadeviceisonlineAPIRequest {
	return &YunosservicecmnscoadeviceisonlineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoadeviceisonlineAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.device.isonline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoadeviceisonlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoadeviceisonlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 设备id类型，取值&#34;uuid&#34;或者&#34;imei&#34;或者&#34;deviceToken&#34;
func (r *YunosservicecmnscoadeviceisonlineAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r YunosservicecmnscoadeviceisonlineAPIRequest) GetType() string {
	return r._type
}

// SetValue is Value Setter
// 对应的设备id值
func (r *YunosservicecmnscoadeviceisonlineAPIRequest) SetValue(_value string) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r YunosservicecmnscoadeviceisonlineAPIRequest) GetValue() string {
	return r._value
}
