package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoservindustryfinanceinsuranceagreementsignAPIRequest 保司合同签约后回调接口 API请求
// taobao.servindustry.finance.insurance.agreement.sign
//
// 保司合同签约后回调接口
type TaobaoservindustryfinanceinsuranceagreementsignAPIRequest struct {
	model.Params
	// 内部模板id
	_innerAgreementId string
	// 唯一码
	_uniqueCode string
	// 保司模板id
	_outAgreementId string
}

// NewTaobaoservindustryfinanceinsuranceagreementsignRequest 初始化TaobaoservindustryfinanceinsuranceagreementsignAPIRequest对象
func NewTaobaoservindustryfinanceinsuranceagreementsignRequest() *TaobaoservindustryfinanceinsuranceagreementsignAPIRequest {
	return &TaobaoservindustryfinanceinsuranceagreementsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.insurance.agreement.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInnerAgreementId is InnerAgreementId Setter
// 内部模板id
func (r *TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) SetInnerAgreementId(_innerAgreementId string) error {
	r._innerAgreementId = _innerAgreementId
	r.Set("inner_agreement_id", _innerAgreementId)
	return nil
}

// GetInnerAgreementId InnerAgreementId Getter
func (r TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) GetInnerAgreementId() string {
	return r._innerAgreementId
}

// SetUniqueCode is UniqueCode Setter
// 唯一码
func (r *TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) SetUniqueCode(_uniqueCode string) error {
	r._uniqueCode = _uniqueCode
	r.Set("unique_code", _uniqueCode)
	return nil
}

// GetUniqueCode UniqueCode Getter
func (r TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) GetUniqueCode() string {
	return r._uniqueCode
}

// SetOutAgreementId is OutAgreementId Setter
// 保司模板id
func (r *TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) SetOutAgreementId(_outAgreementId string) error {
	r._outAgreementId = _outAgreementId
	r.Set("out_agreement_id", _outAgreementId)
	return nil
}

// GetOutAgreementId OutAgreementId Getter
func (r TaobaoservindustryfinanceinsuranceagreementsignAPIRequest) GetOutAgreementId() string {
	return r._outAgreementId
}
