package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopInterceptAPIRequest 阿里巴巴.天猫. 履约订单. 配送拦截 API请求
// alibaba.tianmao.uop.intercept
//
// 阿里巴巴.天猫. 履约订单. 配送拦截
type AlibabaTianmaoUopInterceptAPIRequest struct {
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

// NewAlibabaTianmaoUopInterceptRequest 初始化AlibabaTianmaoUopInterceptAPIRequest对象
func NewAlibabaTianmaoUopInterceptRequest() *AlibabaTianmaoUopInterceptAPIRequest {
	return &AlibabaTianmaoUopInterceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoUopInterceptAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.uop.intercept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoUopInterceptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetScpCode is ScpCode Setter
// SCP单号
func (r *AlibabaTianmaoUopInterceptAPIRequest) SetScpCode(_scpCode string) error {
	r._scpCode = _scpCode
	r.Set("scp_code", _scpCode)
	return nil
}

// GetScpCode ScpCode Getter
func (r AlibabaTianmaoUopInterceptAPIRequest) GetScpCode() string {
	return r._scpCode
}

// SetOperatorNick is OperatorNick Setter
// 操作人名称
func (r *AlibabaTianmaoUopInterceptAPIRequest) SetOperatorNick(_operatorNick string) error {
	r._operatorNick = _operatorNick
	r.Set("operator_nick", _operatorNick)
	return nil
}

// GetOperatorNick OperatorNick Getter
func (r AlibabaTianmaoUopInterceptAPIRequest) GetOperatorNick() string {
	return r._operatorNick
}

// SetMailNos is MailNos Setter
// 运单号
func (r *AlibabaTianmaoUopInterceptAPIRequest) SetMailNos(_mailNos string) error {
	r._mailNos = _mailNos
	r.Set("mail_nos", _mailNos)
	return nil
}

// GetMailNos MailNos Getter
func (r AlibabaTianmaoUopInterceptAPIRequest) GetMailNos() string {
	return r._mailNos
}

// SetOwnerId is OwnerId Setter
// 货主id
func (r *AlibabaTianmaoUopInterceptAPIRequest) SetOwnerId(_ownerId int64) error {
	r._ownerId = _ownerId
	r.Set("owner_id", _ownerId)
	return nil
}

// GetOwnerId OwnerId Getter
func (r AlibabaTianmaoUopInterceptAPIRequest) GetOwnerId() int64 {
	return r._ownerId
}
