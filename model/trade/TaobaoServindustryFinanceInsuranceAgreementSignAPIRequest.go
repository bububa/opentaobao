package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) Reset() {
	r._innerAgreementId = ""
	r._uniqueCode = ""
	r._outAgreementId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.insurance.agreement.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoServindustryFinanceInsuranceAgreementSignAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoServindustryFinanceInsuranceAgreementSignRequest()
	},
}

// GetTaobaoServindustryFinanceInsuranceAgreementSignRequest 从 sync.Pool 获取 TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest
func GetTaobaoServindustryFinanceInsuranceAgreementSignAPIRequest() *TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest {
	return poolTaobaoServindustryFinanceInsuranceAgreementSignAPIRequest.Get().(*TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest)
}

// ReleaseTaobaoServindustryFinanceInsuranceAgreementSignAPIRequest 将 TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest 放入 sync.Pool
func ReleaseTaobaoServindustryFinanceInsuranceAgreementSignAPIRequest(v *TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest) {
	v.Reset()
	poolTaobaoServindustryFinanceInsuranceAgreementSignAPIRequest.Put(v)
}
