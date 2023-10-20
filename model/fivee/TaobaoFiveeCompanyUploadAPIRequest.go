package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveecompanyuploadAPIRequest 上传商信息接口 API请求
// taobao.fivee.company.upload
//
// 资质共享平台上传资质证照
type TaobaofiveecompanyuploadAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 商家证照信息
	_paramCompany *Company
}

// NewTaobaofiveecompanyuploadRequest 初始化TaobaofiveecompanyuploadAPIRequest对象
func NewTaobaofiveecompanyuploadRequest() *TaobaofiveecompanyuploadAPIRequest {
	return &TaobaofiveecompanyuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofiveecompanyuploadAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.company.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofiveecompanyuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofiveecompanyuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaofiveecompanyuploadAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaofiveecompanyuploadAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetParamCompany is ParamCompany Setter
// 商家证照信息
func (r *TaobaofiveecompanyuploadAPIRequest) SetParamCompany(_paramCompany *Company) error {
	r._paramCompany = _paramCompany
	r.Set("param_company", _paramCompany)
	return nil
}

// GetParamCompany ParamCompany Getter
func (r TaobaofiveecompanyuploadAPIRequest) GetParamCompany() *Company {
	return r._paramCompany
}
