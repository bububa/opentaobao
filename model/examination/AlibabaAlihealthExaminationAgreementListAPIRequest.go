package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationAgreementListAPIRequest isv协议获取 API请求
// alibaba.alihealth.examination.agreement.list
//
// isv协议获取
type AlibabaAlihealthExaminationAgreementListAPIRequest struct {
	model.Params
	// isv传递过来的门店code
	_storeCode string
}

// NewAlibabaAlihealthExaminationAgreementListRequest 初始化AlibabaAlihealthExaminationAgreementListAPIRequest对象
func NewAlibabaAlihealthExaminationAgreementListRequest() *AlibabaAlihealthExaminationAgreementListAPIRequest {
	return &AlibabaAlihealthExaminationAgreementListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationAgreementListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.agreement.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationAgreementListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationAgreementListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// isv传递过来的门店code
func (r *AlibabaAlihealthExaminationAgreementListAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r AlibabaAlihealthExaminationAgreementListAPIRequest) GetStoreCode() string {
	return r._storeCode
}
