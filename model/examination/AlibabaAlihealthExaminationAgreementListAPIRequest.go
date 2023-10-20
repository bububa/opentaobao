package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationagreementlistAPIRequest isv协议获取 API请求
// alibaba.alihealth.examination.agreement.list
//
// isv协议获取
type AlibabaalihealthexaminationagreementlistAPIRequest struct {
	model.Params
	// isv传递过来的门店code
	_storeCode string
}

// NewAlibabaalihealthexaminationagreementlistRequest 初始化AlibabaalihealthexaminationagreementlistAPIRequest对象
func NewAlibabaalihealthexaminationagreementlistRequest() *AlibabaalihealthexaminationagreementlistAPIRequest {
	return &AlibabaalihealthexaminationagreementlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationagreementlistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.agreement.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationagreementlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationagreementlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// isv传递过来的门店code
func (r *AlibabaalihealthexaminationagreementlistAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r AlibabaalihealthexaminationagreementlistAPIRequest) GetStoreCode() string {
	return r._storeCode
}
