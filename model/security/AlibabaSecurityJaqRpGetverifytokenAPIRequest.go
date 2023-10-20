package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpGetverifytokenAPIRequest 聚安全实人认证获取认证会话token API请求
// alibaba.security.jaq.rp.getverifytoken
//
// 聚安全实人认证获取认证会话token
type AlibabaSecurityJaqRpGetverifytokenAPIRequest struct {
	model.Params
	// 账号，强烈建议填写，区别用户的唯一标识
	_accountId string
	// 选填，作为一次验证的唯一标识，每次验证需更换。如果不是验证类型可不填
	_ticketId string
	// 客户端来源
	_source string
	// 业务点
	_biz string
	// 额外信息
	_extraData string
}

// NewAlibabaSecurityJaqRpGetverifytokenRequest 初始化AlibabaSecurityJaqRpGetverifytokenAPIRequest对象
func NewAlibabaSecurityJaqRpGetverifytokenRequest() *AlibabaSecurityJaqRpGetverifytokenAPIRequest {
	return &AlibabaSecurityJaqRpGetverifytokenAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpGetverifytokenAPIRequest) Reset() {
	r._accountId = ""
	r._ticketId = ""
	r._source = ""
	r._biz = ""
	r._extraData = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.getverifytoken"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountId is AccountId Setter
// 账号，强烈建议填写，区别用户的唯一标识
func (r *AlibabaSecurityJaqRpGetverifytokenAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetTicketId is TicketId Setter
// 选填，作为一次验证的唯一标识，每次验证需更换。如果不是验证类型可不填
func (r *AlibabaSecurityJaqRpGetverifytokenAPIRequest) SetTicketId(_ticketId string) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetTicketId() string {
	return r._ticketId
}

// SetSource is Source Setter
// 客户端来源
func (r *AlibabaSecurityJaqRpGetverifytokenAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetSource() string {
	return r._source
}

// SetBiz is Biz Setter
// 业务点
func (r *AlibabaSecurityJaqRpGetverifytokenAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetBiz() string {
	return r._biz
}

// SetExtraData is ExtraData Setter
// 额外信息
func (r *AlibabaSecurityJaqRpGetverifytokenAPIRequest) SetExtraData(_extraData string) error {
	r._extraData = _extraData
	r.Set("extra_data", _extraData)
	return nil
}

// GetExtraData ExtraData Getter
func (r AlibabaSecurityJaqRpGetverifytokenAPIRequest) GetExtraData() string {
	return r._extraData
}

var poolAlibabaSecurityJaqRpGetverifytokenAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpGetverifytokenRequest()
	},
}

// GetAlibabaSecurityJaqRpGetverifytokenRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpGetverifytokenAPIRequest
func GetAlibabaSecurityJaqRpGetverifytokenAPIRequest() *AlibabaSecurityJaqRpGetverifytokenAPIRequest {
	return poolAlibabaSecurityJaqRpGetverifytokenAPIRequest.Get().(*AlibabaSecurityJaqRpGetverifytokenAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpGetverifytokenAPIRequest 将 AlibabaSecurityJaqRpGetverifytokenAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpGetverifytokenAPIRequest(v *AlibabaSecurityJaqRpGetverifytokenAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpGetverifytokenAPIRequest.Put(v)
}
