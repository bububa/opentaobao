package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLatourStrategyIssueAPIRequest
阿里巴巴权益发放接口 API请求
alibaba.latour.strategy.issue

阿里巴巴权益平台权益发放接口 */
type AlibabaLatourStrategyIssueAPIRequest struct {
	model.Params
	// 扩展参数
	_extraData string
	// 算法容灾
	_failoverAlgorithmResult bool
	// 幂等id
	_idempotentId string
	// 发放渠道
	_channel string
	// 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
	_userId string
	// 转换用户类型
	_transformedUserType string
	// 是否需要过安全
	_needIdentifyRisk bool
	// 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
	_userNick string
	// 投放计划code
	_strategyCode string
	// 用户类型
	_userType string
	// 指定发放权益code
	_selectedBenefitCode string
}

// New
