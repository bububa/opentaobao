package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaouopinterceptAPIRequest 阿里巴巴.天猫. 履约订单. 配送拦截 API请求
// alibaba.tianmao.uop.intercept
//
// 阿里巴巴.天猫. 履约订单. 配送拦截
type AlibabatianmaouopinterceptAPIRequest struct {
	model.Params
	// SCP单号
	_scpCode string
	// 操作人名称
	_operatorNick string
	// 运单号
	_mailNos string
	// 货主id
	_ownerId int64
}

// NewAlibabatianmaouopinterceptRequest 初始化AlibabatianmaouopinterceptAPIRequest对象
func NewAlibabatianmaouopinterceptRequest() *AlibabatianmaouopinterceptAPIRequest {
	return &AlibabatianmaouopinterceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianmaouopinterceptAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.uop.intercept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianmaouopinterceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianmaouopinterceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScpCode is ScpCode Setter
// SCP单号
func (r *AlibabatianmaouopinterceptAPIRequest) SetScpCode(_scpCode string) error {
	r._scpCode = _scpCode
	r.Set("scp_code", _scpCode)
	return nil
}

// GetScpCode ScpCode Getter
func (r AlibabatianmaouopinterceptAPIRequest) GetScpCode() string {
	return r._scpCode
}

// SetOperatorNick is OperatorNick Setter
// 操作人名称
func (r *AlibabatianmaouopinterceptAPIRequest) SetOperatorNick(_operatorNick string) error {
	r._operatorNick = _operatorNick
	r.Set("operator_nick", _operatorNick)
	return nil
}

// GetOperatorNick OperatorNick Getter
func (r AlibabatianmaouopinterceptAPIRequest) GetOperatorNick() string {
	return r._operatorNick
}

// SetMailNos is MailNos Setter
// 运单号
func (r *AlibabatianmaouopinterceptAPIRequest) SetMailNos(_mailNos string) error {
	r._mailNos = _mailNos
	r.Set("mail_nos", _mailNos)
	return nil
}

// GetMailNos MailNos Getter
func (r AlibabatianmaouopinterceptAPIRequest) GetMailNos() string {
	return r._mailNos
}

// SetOwnerId is OwnerId Setter
// 货主id
func (r *AlibabatianmaouopinterceptAPIRequest) SetOwnerId(_ownerId int64) error {
	r._ownerId = _ownerId
	r.Set("owner_id", _ownerId)
	return nil
}

// GetOwnerId OwnerId Getter
func (r AlibabatianmaouopinterceptAPIRequest) GetOwnerId() int64 {
	return r._ownerId
}
