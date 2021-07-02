package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLatourStrategyShowAPIRequest 阿里巴巴权益投放接口 API请求
// alibaba.latour.strategy.show
//
// 阿里巴巴权益平台权益投放接口
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

// NewAlibabaLatourStrategyShowRequest 初始化AlibabaLatourStrategyShowAPIRequest对象
func NewAlibabaLatourStrategyShowRequest() *AlibabaLatourStrategyShowAPIRequest {
	return &AlibabaLatourStrategyShowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLatourStrategyShowAPIRequest) GetApiMethodName() string {
	return "alibaba.latour.strategy.show"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLatourStrategyShowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WithTestBenefit Setter
// 带出测试权益
func (r *AlibabaLatourStrategyShowAPIRequest) SetWithTestBenefit(_withTestBenefit bool) error {
	r._withTestBenefit = _withTestBenefit
	r.Set("with_test_benefit", _withTestBenefit)
	return nil
}

// Get WithTestBenefit Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetWithTestBenefit() bool {
	return r._withTestBenefit
}

// Set is Channel Setter
// 渠道
func (r *AlibabaLatourStrategyShowAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetChannel() string {
	return r._channel
}

// Set is PageSize Setter
// 每页权益数
func (r *AlibabaLatourStrategyShowAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is TransformedUserType Setter
// 要转换的账户类型
func (r *AlibabaLatourStrategyShowAPIRequest) SetTransformedUserType(_transformedUserType string) error {
	r._transformedUserType = _transformedUserType
	r.Set("transformed_user_type", _transformedUserType)
	return nil
}

// Get TransformedUserType Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetTransformedUserType() string {
	return r._transformedUserType
}

// Set is NeedIdentifyRisk Setter
// 是否需要调用安全校验服务
func (r *AlibabaLatourStrategyShowAPIRequest) SetNeedIdentifyRisk(_needIdentifyRisk bool) error {
	r._needIdentifyRisk = _needIdentifyRisk
	r.Set("need_identify_risk", _needIdentifyRisk)
	return nil
}

// Get NeedIdentifyRisk Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetNeedIdentifyRisk() bool {
	return r._needIdentifyRisk
}

// Set is UserNick Setter
// 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyShowAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// Get UserNick Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetUserNick() string {
	return r._userNick
}

// Set is SkipWithHadWin Setter
// 不带出hadWin状态
func (r *AlibabaLatourStrategyShowAPIRequest) SetSkipWithHadWin(_skipWithHadWin bool) error {
	r._skipWithHadWin = _skipWithHadWin
	r.Set("skip_with_had_win", _skipWithHadWin)
	return nil
}

// Get SkipWithHadWin Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetSkipWithHadWin() bool {
	return r._skipWithHadWin
}

// Set is FilterEmptyInventory Setter
// 过滤无库存权益
func (r *AlibabaLatourStrategyShowAPIRequest) SetFilterEmptyInventory(_filterEmptyInventory bool) error {
	r._filterEmptyInventory = _filterEmptyInventory
	r.Set("filter_empty_inventory", _filterEmptyInventory)
	return nil
}

// Get FilterEmptyInventory Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetFilterEmptyInventory() bool {
	return r._filterEmptyInventory
}

// Set is UserId Setter
// 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyShowAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetUserId() string {
	return r._userId
}

// Set is StrategyCode Setter
// 投放计划code
func (r *AlibabaLatourStrategyShowAPIRequest) SetStrategyCode(_strategyCode string) error {
	r._strategyCode = _strategyCode
	r.Set("strategy_code", _strategyCode)
	return nil
}

// Get StrategyCode Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetStrategyCode() string {
	return r._strategyCode
}

// Set is UserType Setter
// 当面账户类型
func (r *AlibabaLatourStrategyShowAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// Get UserType Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetUserType() string {
	return r._userType
}

// Set is CurrentPage Setter
// 当面分页
func (r *AlibabaLatourStrategyShowAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// Set is FilterCrowd Setter
// 过滤人群
func (r *AlibabaLatourStrategyShowAPIRequest) SetFilterCrowd(_filterCrowd bool) error {
	r._filterCrowd = _filterCrowd
	r.Set("filter_crowd", _filterCrowd)
	return nil
}

// Get FilterCrowd Getter
func (r AlibabaLatourStrategyShowAPIRequest) GetFilterCrowd() bool {
	return r._filterCrowd
}
