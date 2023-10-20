package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosservicecmnscoadevicegetAPIRequest 设备详情查询 API请求
// yunos.service.cmns.coa.device.get
//
// 第三方应用开发者调用此接口查询设备详情
type YunosservicecmnscoadevicegetAPIRequest struct {
	model.Params
	// 设备id类型,可以是uuid,imei,deviceToken,kp
	_type string
	// 设备id
	_value string
}

// NewYunosservicecmnscoadevicegetRequest 初始化YunosservicecmnscoadevicegetAPIRequest对象
func NewYunosservicecmnscoadevicegetRequest() *YunosservicecmnscoadevicegetAPIRequest {
	return &YunosservicecmnscoadevicegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosservicecmnscoadevicegetAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.device.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosservicecmnscoadevicegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosservicecmnscoadevicegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 设备id类型,可以是uuid,imei,deviceToken,kp
func (r *YunosservicecmnscoadevicegetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r YunosservicecmnscoadevicegetAPIRequest) GetType() string {
	return r._type
}

// SetValue is Value Setter
// 设备id
func (r *YunosservicecmnscoadevicegetAPIRequest) SetValue(_value string) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r YunosservicecmnscoadevicegetAPIRequest) GetValue() string {
	return r._value
}
