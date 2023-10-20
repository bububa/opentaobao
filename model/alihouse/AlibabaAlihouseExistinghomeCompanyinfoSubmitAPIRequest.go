package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest 进件接口 API请求
// alibaba.alihouse.existinghome.companyinfo.submit
//
// 进件接口
type AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest struct {
	model.Params
	// 进件资料对象
	_cis *TradeMerchantOpenDto
}

// NewAlibabaalihouseexistinghomecompanyinfosubmitRequest 初始化AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest对象
func NewAlibabaalihouseexistinghomecompanyinfosubmitRequest() *AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest {
	return &AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.companyinfo.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCis is Cis Setter
// 进件资料对象
func (r *AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest) SetCis(_cis *TradeMerchantOpenDto) error {
	r._cis = _cis
	r.Set("cis", _cis)
	return nil
}

// GetCis Cis Getter
func (r AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest) GetCis() *TradeMerchantOpenDto {
	return r._cis
}
