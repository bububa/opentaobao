package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2cmemberdiscountrulegetAPIRequest 会员抵扣规则 API请求
// alibaba.nlife.b2c.member.discountrule.get
//
// 获取企业会员抵扣规则
type Alibabanlifeb2cmemberdiscountrulegetAPIRequest struct {
	model.Params
	// 企业ID
	_companyId string
	// 会员在ISV处的编号
	_cardNo string
}

// NewAlibabanlifeb2cmemberdiscountrulegetRequest 初始化Alibabanlifeb2cmemberdiscountrulegetAPIRequest对象
func NewAlibabanlifeb2cmemberdiscountrulegetRequest() *Alibabanlifeb2cmemberdiscountrulegetAPIRequest {
	return &Alibabanlifeb2cmemberdiscountrulegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r Alibabanlifeb2cmemberdiscountrulegetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.member.discountrule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r Alibabanlifeb2cmemberdiscountrulegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r Alibabanlifeb2cmemberdiscountrulegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyId is CompanyId Setter
// 企业ID
func (r *Alibabanlifeb2cmemberdiscountrulegetAPIRequest) SetCompanyId(_companyId string) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r Alibabanlifeb2cmemberdiscountrulegetAPIRequest) GetCompanyId() string {
	return r._companyId
}

// SetCardNo is CardNo Setter
// 会员在ISV处的编号
func (r *Alibabanlifeb2cmemberdiscountrulegetAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r Alibabanlifeb2cmemberdiscountrulegetAPIRequest) GetCardNo() string {
	return r._cardNo
}
