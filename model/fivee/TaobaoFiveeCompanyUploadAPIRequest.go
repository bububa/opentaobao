package fivee

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFiveeCompanyUploadAPIRequest) Reset() {
	r._paramBucode = ""
	r._paramCompany = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFiveeCompanyUploadAPIRequest) GetApiMethodName() string {
	return "taobao.fivee.company.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFiveeCompanyUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFiveeCompanyUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBucode is ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeCompanyUploadAPIRequest) SetParamBucode(_paramBucode string) error {
	r._paramBucode = _paramBucode
	r.Set("param_bucode", _paramBucode)
	return nil
}

// GetParamBucode ParamBucode Getter
func (r TaobaoFiveeCompanyUploadAPIRequest) GetParamBucode() string {
	return r._paramBucode
}

// SetParamCompany is ParamCompany Setter
// 商家证照信息
func (r *TaobaoFiveeCompanyUploadAPIRequest) SetParamCompany(_paramCompany *Company) error {
	r._paramCompany = _paramCompany
	r.Set("param_company", _paramCompany)
	return nil
}

// GetParamCompany ParamCompany Getter
func (r TaobaoFiveeCompanyUploadAPIRequest) GetParamCompany() *Company {
	return r._paramCompany
}

var poolTaobaoFiveeCompanyUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFiveeCompanyUploadRequest()
	},
}

// GetTaobaoFiveeCompanyUploadRequest 从 sync.Pool 获取 TaobaoFiveeCompanyUploadAPIRequest
func GetTaobaoFiveeCompanyUploadAPIRequest() *TaobaoFiveeCompanyUploadAPIRequest {
	return poolTaobaoFiveeCompanyUploadAPIRequest.Get().(*TaobaoFiveeCompanyUploadAPIRequest)
}

// ReleaseTaobaoFiveeCompanyUploadAPIRequest 将 TaobaoFiveeCompanyUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoFiveeCompanyUploadAPIRequest(v *TaobaoFiveeCompanyUploadAPIRequest) {
	v.Reset()
	poolTaobaoFiveeCompanyUploadAPIRequest.Put(v)
}
