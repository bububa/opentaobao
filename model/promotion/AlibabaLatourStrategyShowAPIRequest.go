package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLatourStrategyShowAPIRequest
阿里巴巴权益投放接口 API请求
alibaba.latour.strategy.show

阿里巴巴权益平台权益投放接口 */
type AlibabaLatourStrategyShowAPIRequest struct {
	model.Params
	// 带出测试权益
	_withTestBenefit bool
	// 渠道
	_channel string
	// 每页权益数
	_pageSize int64
	// 要转换的账户类型
	_transformedUserType string
	// 是否需要调用安全校验服务
	_needIdentifyRisk bool
	// 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
	_userNick string
	// 不带出hadWin状态
	_skipWithHadWin bool
	// 过滤无库存权益
	_filterEmptyInventory bool
	// 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
	_userId string
	// 投放计划code
	_strategyCode string
	// 当面账户类型
	_userType string
	// 当面分页
	_currentPage int64
	// 过滤人群
	_filterCrowd bool
}

// New
