package cainiaoecc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEccExceptionsDelayGetAPIRequest
菜鸟控制塔包裹滞留异常信息获取 API请求
cainiao.ecc.exceptions.delay.get

菜鸟控制塔包裹滞留异常信息获取 */
type CainiaoEccExceptionsDelayGetAPIRequest struct {
	model.Params
	// 运单号
	_mailNo string
}

// NewCainiaoEccExceptionsDelayGetRequest 初始化CainiaoEccExceptionsDelayGetAPIRequest对象
func NewCainiaoEccExceptionsDelayGetRequest() *CainiaoEccExceptionsDelayGetAPIRequest {
	return &CainiaoEccExceptionsDelayGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEccExceptionsDelayGetAPIRequest) GetApiMethodName() string {
	return "cainiao.ecc.exceptions.delay.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEccExceptionsDelayGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MailNo Setter
// 运单号
func (r *CainiaoEccExceptionsDelayGetAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// Get MailNo Getter
func (r CainiaoEccExceptionsDelayGetAPIRequest) GetMailNo() string {
	return r._mailNo
}
