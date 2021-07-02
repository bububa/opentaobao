package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeCompanyGetAPIRequest 查询商信息 API请求
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
type TaobaoFiveeCompanyGetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 统一社会信息用代码
	_paramUniqueCode string
}

// NewTaobaoFiveeCompanyGetRequest 初始化TaobaoFiveeCompanyGetAPIRequest对象
func NewTaobaoFiveeCompanyGetRequest() *TaobaoFiveeCompanyGetAPIRequest {
	return &TaobaoFiveeCompanyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeCompanyGetAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.company.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeCompanyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeCompanyGetAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaoFiveeCompanyGetAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetParamUniqueCode is ParamUniqueCode Setter
// 统一社会信息用代码
func (r *TaobaoFiveeCompanyGetAPIRequest) SetParamUniqueCode(_paramUniqueCode string) error {
	r._paramUniqueCode = _paramUniqueCode
	r.Set("param_unique_code", _paramUniqueCode)
	return nil
}

// GetParamUniqueCode ParamUniqueCode Getter
func (r TaobaoFiveeCompanyGetAPIRequest) GetParamUniqueCode() string {
	return r._paramUniqueCode
}
