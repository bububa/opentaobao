package cainiaoecc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoeccexceptionsdelaygetAPIRequest 菜鸟控制塔包裹滞留异常信息获取 API请求
// cainiao.ecc.exceptions.delay.get
//
// 菜鸟控制塔包裹滞留异常信息获取
type CainiaoeccexceptionsdelaygetAPIRequest struct {
	model.Params
	// 运单号
	_mailNo string
}

// NewCainiaoeccexceptionsdelaygetRequest 初始化CainiaoeccexceptionsdelaygetAPIRequest对象
func NewCainiaoeccexceptionsdelaygetRequest() *CainiaoeccexceptionsdelaygetAPIRequest {
	return &CainiaoeccexceptionsdelaygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoeccexceptionsdelaygetAPIRequest) GetApiMethodName() string {
	return "cainiao.ecc.exceptions.delay.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoeccexceptionsdelaygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoeccexceptionsdelaygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoeccexceptionsdelaygetAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoeccexceptionsdelaygetAPIRequest) GetMailNo() string {
	return r._mailNo
}
