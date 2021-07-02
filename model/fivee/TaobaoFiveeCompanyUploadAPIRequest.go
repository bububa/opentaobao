package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeCompanyUploadAPIRequest 上传商信息接口 API请求
// taobao.fivee.company.upload
//
// 资质共享平台上传资质证照
type TaobaoFiveeCompanyUploadAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 商家证照信息
	_paramCompany *Company
}

// NewTaobaoFiveeCompanyUploadRequest 初始化TaobaoFiveeCompanyUploadAPIRequest对象
func NewTaobaoFiveeCompanyUploadRequest() *TaobaoFiveeCompanyUploadAPIRequest {
	return &TaobaoFiveeCompanyUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeCompanyUploadAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.company.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeCompanyUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeCompanyUploadAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// Get ParamBucode Getter
func (r TaobaoFiveeCompanyUploadAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// Set is ParamCompany Setter
// 商家证照信息
func (r *TaobaoFiveeCompanyUploadAPIRequest) SetParamCompany(_paramCompany *Company) error {
	r._paramCompany = _paramCompany
	r.Set("param_company", _paramCompany)
	return nil
}

// Get ParamCompany Getter
func (r TaobaoFiveeCompanyUploadAPIRequest) GetParamCompany() *Company {
	return r._paramCompany
}
