package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerEntSearchAPIRequest 查询商家信息 API请求
// alibaba.alihealth.tracecodeseller.ent.search
//
// 查询商家信息
type AlibabaAlihealthTracecodesellerEntSearchAPIRequest struct {
	model.Params
	// appkey
	_skeyCode string
	// 商家名称
	_name string
	// 淘宝名
	_tbUserId string
}

// NewAlibabaAlihealthTracecodesellerEntSearchRequest 初始化AlibabaAlihealthTracecodesellerEntSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerEntSearchRequest() *AlibabaAlihealthTracecodesellerEntSearchAPIRequest {
	return &AlibabaAlihealthTracecodesellerEntSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerEntSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.ent.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerEntSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkeyCode Setter
// appkey
func (r *AlibabaAlihealthTracecodesellerEntSearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// Get SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerEntSearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// Set is Name Setter
// 商家名称
func (r *AlibabaAlihealthTracecodesellerEntSearchAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r AlibabaAlihealthTracecodesellerEntSearchAPIRequest) GetName() string {
	return r._name
}

// Set is TbUserId Setter
// 淘宝名
func (r *AlibabaAlihealthTracecodesellerEntSearchAPIRequest) SetTbUserId(_tbUserId string) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// Get TbUserId Getter
func (r AlibabaAlihealthTracecodesellerEntSearchAPIRequest) GetTbUserId() string {
	return r._tbUserId
}
