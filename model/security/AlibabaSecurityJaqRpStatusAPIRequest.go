package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpstatusAPIRequest 聚安全实人认证查询状态接口 API请求
// alibaba.security.jaq.rp.status
//
// 聚安全实人认证查询状态接口
type AlibabasecurityjaqrpstatusAPIRequest struct {
	model.Params
	// 账号id
	_accountId string
	// 凭据id
	_ticketId string
	// 客户端来源
	_source string
	// 业务来源
	_biz string
}

// NewAlibabasecurityjaqrpstatusRequest 初始化AlibabasecurityjaqrpstatusAPIRequest对象
func NewAlibabasecurityjaqrpstatusRequest() *AlibabasecurityjaqrpstatusAPIRequest {
	return &AlibabasecurityjaqrpstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 账号id
func (r *AlibabasecurityjaqrpstatusAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabasecurityjaqrpstatusAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetTicketId is TicketId Setter
// 凭据id
func (r *AlibabasecurityjaqrpstatusAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r AlibabasecurityjaqrpstatusAPIRequest) GetTicketId() string {
	return r._ticketId
}

// SetSource is Source Setter
// 客户端来源
func (r *AlibabasecurityjaqrpstatusAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabasecurityjaqrpstatusAPIRequest) GetSource() string {
	return r._source
}

// SetBiz is Biz Setter
// 业务来源
func (r *AlibabasecurityjaqrpstatusAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r AlibabasecurityjaqrpstatusAPIRequest) GetBiz() string {
	return r._biz
}
