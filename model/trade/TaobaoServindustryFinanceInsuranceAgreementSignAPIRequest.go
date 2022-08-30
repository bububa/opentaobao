package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest 保司合同签约后回调接口 API请求
// taobao.servindustry.finance.insurance.agreement.sign
//
// 保司合同签约后回调接口
type TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest struct {
	model.Params
	// 内部模板id
	_innerAgreementId string
	// 唯一码
	_uniqueCode string
	// 保司模板id
	_outAgreementId string
}

// NewTaobaoServindustryFinanceInsuranceAgreementSignRequest 初始化TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest对象
func NewTaobaoServindustryFinanceInsuranceAgreementSignRequest() *TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest {
	return &TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.insurance.agreement.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInnerAgreementId is InnerAgreementId Setter
// 内部模板id
func (r *TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) SetInnerAgreementId(_innerAgreementId string) error {
	r._innerAgreementId = _innerAgreementId
	r.Set("inner_agreement_id", _innerAgreementId)
	return nil
}

// GetInnerAgreementId InnerAgreementId Getter
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetInnerAgreementId() string {
	return r._innerAgreementId
}

// SetUniqueCode is UniqueCode Setter
// 唯一码
func (r *TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) SetUniqueCode(_uniqueCode string) error {
	r._uniqueCode = _uniqueCode
	r.Set("unique_code", _uniqueCode)
	return nil
}

// GetUniqueCode UniqueCode Getter
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetUniqueCode() string {
	return r._uniqueCode
}

// SetOutAgreementId is OutAgreementId Setter
// 保司模板id
func (r *TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) SetOutAgreementId(_outAgreementId string) error {
	r._outAgreementId = _outAgreementId
	r.Set("out_agreement_id", _outAgreementId)
	return nil
}

// GetOutAgreementId OutAgreementId Getter
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetOutAgreementId() string {
	return r._outAgreementId
}
