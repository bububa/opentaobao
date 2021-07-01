package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest
保护结果反馈 API请求
alibaba.security.jaq.spamregisterprevention.feedback

用户通过这个接口对垃圾注册防控结果进行反馈 */
type AlibabaSecurityJaqSpamregisterpreventionFeedbackAPIRequest struct {
	model.Params
	// 查询接口返回的id
	_itemId string
	// 聚安全返回的决定信息。当feedBack为0时可以不添
	_jaqDecision int64
	// 用户自己的决定信息。当feedBack为0时可以不添。
	_customerDecision int64
	// 用户不认同聚安全返回结果的原因类型。0：同意判定；1：和三方结果不符；2：用户投诉；3:经过人工review判断;9:	其他。
	_feedBack int64
	// 用户不认同聚安全返回结果的原因描述
	_denyReason string
}

// New
