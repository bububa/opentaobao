package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest 保护结果反馈 API请求
// alibaba.security.jaq.spamregisterprevention.feedback
//
// 用户通过这个接口对垃圾注册防控结果进行反馈
type AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest struct {
	model.Params
	// 查询接口返回的id
	_itemId string
	// 用户不认同聚安全返回结果的原因描述
	_denyReason string
	// 聚安全返回的决定信息。当feedBack为0时可以不添
	_jaqDecision int64
	// 用户自己的决定信息。当feedBack为0时可以不添。
	_customerDecision int64
	// 用户不认同聚安全返回结果的原因类型。0：同意判定；1：和三方结果不符；2：用户投诉；3:经过人工review判断;9:	其他。
	_feedBack int64
}

// NewAlibabaSecurityJaqSpamregisterpreventionFeedbackRequest 初始化AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest对象
func NewAlibabaSecurityJaqSpamregisterpreventionFeedbackRequest() *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest {
	return &AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) Reset() {
	r._itemId = ""
	r._denyReason = ""
	r._jaqDecision = 0
	r._customerDecision = 0
	r._feedBack = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.spamregisterprevention.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 查询接口返回的id
func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetItemId() string {
	return r._itemId
}

// SetDenyReason is DenyReason Setter
// 用户不认同聚安全返回结果的原因描述
func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) SetDenyReason(_denyReason string) error {
	r._denyReason = _denyReason
	r.Set("deny_reason", _denyReason)
	return nil
}

// GetDenyReason DenyReason Getter
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetDenyReason() string {
	return r._denyReason
}

// SetJaqDecision is JaqDecision Setter
// 聚安全返回的决定信息。当feedBack为0时可以不添
func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) SetJaqDecision(_jaqDecision int64) error {
	r._jaqDecision = _jaqDecision
	r.Set("jaq_decision", _jaqDecision)
	return nil
}

// GetJaqDecision JaqDecision Getter
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetJaqDecision() int64 {
	return r._jaqDecision
}

// SetCustomerDecision is CustomerDecision Setter
// 用户自己的决定信息。当feedBack为0时可以不添。
func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) SetCustomerDecision(_customerDecision int64) error {
	r._customerDecision = _customerDecision
	r.Set("customer_decision", _customerDecision)
	return nil
}

// GetCustomerDecision CustomerDecision Getter
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetCustomerDecision() int64 {
	return r._customerDecision
}

// SetFeedBack is FeedBack Setter
// 用户不认同聚安全返回结果的原因类型。0：同意判定；1：和三方结果不符；2：用户投诉；3:经过人工review判断;9:	其他。
func (r *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) SetFeedBack(_feedBack int64) error {
	r._feedBack = _feedBack
	r.Set("feed_back", _feedBack)
	return nil
}

// GetFeedBack FeedBack Getter
func (r AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) GetFeedBack() int64 {
	return r._feedBack
}

var poolAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqSpamregisterpreventionFeedbackRequest()
	},
}

// GetAlibabaSecurityJaqSpamregisterpreventionFeedbackRequest 从 sync.Pool 获取 AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest
func GetAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest() *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest {
	return poolAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest.Get().(*AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest)
}

// ReleaseAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest 将 AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest(v *AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest.Put(v)
}
