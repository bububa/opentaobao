package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpStatusAPIRequest 聚安全实人认证查询状态接口 API请求
// alibaba.security.jaq.rp.status
//
// 聚安全实人认证查询状态接口
type AlibabaSecurityJaqRpStatusAPIRequest struct {
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

// NewAlibabaSecurityJaqRpStatusRequest 初始化AlibabaSecurityJaqRpStatusAPIRequest对象
func NewAlibabaSecurityJaqRpStatusRequest() *AlibabaSecurityJaqRpStatusAPIRequest {
	return &AlibabaSecurityJaqRpStatusAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpStatusAPIRequest) Reset() {
	r._accountId = ""
	r._ticketId = ""
	r._source = ""
	r._biz = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 账号id
func (r *AlibabaSecurityJaqRpStatusAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaSecurityJaqRpStatusAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetTicketId is TicketId Setter
// 凭据id
func (r *AlibabaSecurityJaqRpStatusAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r AlibabaSecurityJaqRpStatusAPIRequest) GetTicketId() string {
	return r._ticketId
}

// SetSource is Source Setter
// 客户端来源
func (r *AlibabaSecurityJaqRpStatusAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaSecurityJaqRpStatusAPIRequest) GetSource() string {
	return r._source
}

// SetBiz is Biz Setter
// 业务来源
func (r *AlibabaSecurityJaqRpStatusAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r AlibabaSecurityJaqRpStatusAPIRequest) GetBiz() string {
	return r._biz
}

var poolAlibabaSecurityJaqRpStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpStatusRequest()
	},
}

// GetAlibabaSecurityJaqRpStatusRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpStatusAPIRequest
func GetAlibabaSecurityJaqRpStatusAPIRequest() *AlibabaSecurityJaqRpStatusAPIRequest {
	return poolAlibabaSecurityJaqRpStatusAPIRequest.Get().(*AlibabaSecurityJaqRpStatusAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpStatusAPIRequest 将 AlibabaSecurityJaqRpStatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpStatusAPIRequest(v *AlibabaSecurityJaqRpStatusAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpStatusAPIRequest.Put(v)
}
