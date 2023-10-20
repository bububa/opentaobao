package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerentsearchAPIRequest 查询商家信息 API请求
// alibaba.alihealth.tracecodeseller.ent.search
//
// 查询商家信息
type AlibabaalihealthtracecodesellerentsearchAPIRequest struct {
	model.Params
	// appkey
	_skeyCode string
	// 商家名称
	_name string
	// 淘宝名
	_tbUserId string
}

// NewAlibabaalihealthtracecodesellerentsearchRequest 初始化AlibabaalihealthtracecodesellerentsearchAPIRequest对象
func NewAlibabaalihealthtracecodesellerentsearchRequest() *AlibabaalihealthtracecodesellerentsearchAPIRequest {
	return &AlibabaalihealthtracecodesellerentsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellerentsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.ent.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellerentsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellerentsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkeyCode is SkeyCode Setter
// appkey
func (r *AlibabaalihealthtracecodesellerentsearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// GetSkeyCode SkeyCode Getter
func (r AlibabaalihealthtracecodesellerentsearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// SetName is Name Setter
// 商家名称
func (r *AlibabaalihealthtracecodesellerentsearchAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaalihealthtracecodesellerentsearchAPIRequest) GetName() string {
	return r._name
}

// SetTbUserId is TbUserId Setter
// 淘宝名
func (r *AlibabaalihealthtracecodesellerentsearchAPIRequest) SetTbUserId(_tbUserId string) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// GetTbUserId TbUserId Getter
func (r AlibabaalihealthtracecodesellerentsearchAPIRequest) GetTbUserId() string {
	return r._tbUserId
}
