package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest 保护结果反馈 API请求
// alibaba.security.jaq.spamregisterprevention.feedback
//
// 用户通过这个接口对垃圾注册防控结果进行反馈
type AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest struct {
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

// NewAlibabasecurityjaqspamregisterpreventionfeedbackRequest 初始化AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest对象
func NewAlibabasecurityjaqspamregisterpreventionfeedbackRequest() *AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest {
	return &AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.spamregisterprevention.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 查询接口返回的id
func (r *AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetItemId() string {
	return r._itemId
}

// SetDenyReason is DenyReason Setter
// 用户不认同聚安全返回结果的原因描述
func (r *AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) SetDenyReason(_denyReason string) error {
	r._denyReason = _denyReason
	r.Set("deny_reason", _denyReason)
	return nil
}

// GetDenyReason DenyReason Getter
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetDenyReason() string {
	return r._denyReason
}

// SetJaqDecision is JaqDecision Setter
// 聚安全返回的决定信息。当feedBack为0时可以不添
func (r *AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) SetJaqDecision(_jaqDecision int64) error {
	r._jaqDecision = _jaqDecision
	r.Set("jaq_decision", _jaqDecision)
	return nil
}

// GetJaqDecision JaqDecision Getter
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetJaqDecision() int64 {
	return r._jaqDecision
}

// SetCustomerDecision is CustomerDecision Setter
// 用户自己的决定信息。当feedBack为0时可以不添。
func (r *AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) SetCustomerDecision(_customerDecision int64) error {
	r._customerDecision = _customerDecision
	r.Set("customer_decision", _customerDecision)
	return nil
}

// GetCustomerDecision CustomerDecision Getter
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetCustomerDecision() int64 {
	return r._customerDecision
}

// SetFeedBack is FeedBack Setter
// 用户不认同聚安全返回结果的原因类型。0：同意判定；1：和三方结果不符；2：用户投诉；3:经过人工review判断;9:	其他。
func (r *AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) SetFeedBack(_feedBack int64) error {
	r._feedBack = _feedBack
	r.Set("feed_back", _feedBack)
	return nil
}

// GetFeedBack FeedBack Getter
func (r AlibabasecurityjaqspamregisterpreventionfeedbackAPIRequest) GetFeedBack() int64 {
	return r._feedBack
}
