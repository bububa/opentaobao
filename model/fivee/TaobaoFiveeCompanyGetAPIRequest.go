package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveecompanygetAPIRequest 查询商信息 API请求
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
type TaobaofiveecompanygetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 统一社会信息用代码
	_paramUniqueCode string
}

// NewTaobaofiveecompanygetRequest 初始化TaobaofiveecompanygetAPIRequest对象
func NewTaobaofiveecompanygetRequest() *TaobaofiveecompanygetAPIRequest {
	return &TaobaofiveecompanygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofiveecompanygetAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.company.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofiveecompanygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofiveecompanygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaofiveecompanygetAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaofiveecompanygetAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetParamUniqueCode is ParamUniqueCode Setter
// 统一社会信息用代码
func (r *TaobaofiveecompanygetAPIRequest) SetParamUniqueCode(_paramUniqueCode string) error {
	r._paramUniqueCode = _paramUniqueCode
	r.Set("param_unique_code", _paramUniqueCode)
	return nil
}

// GetParamUniqueCode ParamUniqueCode Getter
func (r TaobaofiveecompanygetAPIRequest) GetParamUniqueCode() string {
	return r._paramUniqueCode
}
