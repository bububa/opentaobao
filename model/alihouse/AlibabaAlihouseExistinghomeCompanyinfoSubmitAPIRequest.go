package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest 进件接口 API请求
// alibaba.alihouse.existinghome.companyinfo.submit
//
// 进件接口
type AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest struct {
	model.Params
	// 进件资料对象
	_cis *TradeMerchantOpenDto
}

// NewAlibabaAlihouseExistinghomeCompanyinfoSubmitRequest 初始化AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest对象
func NewAlibabaAlihouseExistinghomeCompanyinfoSubmitRequest() *AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest {
	return &AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.companyinfo.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCis is Cis Setter
// 进件资料对象
func (r *AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest) SetCis(_cis *TradeMerchantOpenDto) error {
	r._cis = _cis
	r.Set("cis", _cis)
	return nil
}

// GetCis Cis Getter
func (r AlibabaAlihouseExistinghomeCompanyinfoSubmitAPIRequest) GetCis() *TradeMerchantOpenDto {
	return r._cis
}
