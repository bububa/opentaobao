package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalatourstrategyshowAPIRequest 阿里巴巴权益投放接口 API请求
// alibaba.latour.strategy.show
//
// 阿里巴巴权益平台权益投放接口
type AlibabalatourstrategyshowAPIRequest struct {
	model.Params
	// 渠道
	_channel string
	// 要转换的账户类型
	_transformedUserType string
	// 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
	_userNick string
	// 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
	_userId string
	// 投放计划code
	_strategyCode string
	// 当面账户类型
	_userType string
	// openid
	_openid string
	// 每页权益数
	_pageSize int64
	// 当面分页
	_currentPage int64
	// 带出测试权益
	_withTestBenefit bool
	// 是否需要调用安全校验服务
	_needIdentifyRisk bool
	// 不带出hadWin状态
	_skipWithHadWin bool
	// 过滤无库存权益
	_filterEmptyInventory bool
	// 过滤人群
	_filterCrowd bool
	// 是否需要投放计划维度的权益核销实例
	_withStrategyInstance bool
	// 是否需要权益维度的权益核销实例
	_withBenefitInstance bool
}

// NewAlibabalatourstrategyshowRequest 初始化AlibabalatourstrategyshowAPIRequest对象
func NewAlibabalatourstrategyshowRequest() *AlibabalatourstrategyshowAPIRequest {
	return &AlibabalatourstrategyshowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalatourstrategyshowAPIRequest) GetApiMethodName() string {
	return "alibaba.latour.strategy.show"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalatourstrategyshowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalatourstrategyshowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannel is Channel Setter
// 渠道
func (r *AlibabalatourstrategyshowAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabalatourstrategyshowAPIRequest) GetChannel() string {
	return r._channel
}

// SetTransformedUserType is TransformedUserType Setter
// 要转换的账户类型
func (r *AlibabalatourstrategyshowAPIRequest) SetTransformedUserType(_transformedUserType string) error {
	r._transformedUserType = _transformedUserType
	r.Set("transformed_user_type", _transformedUserType)
	return nil
}

// GetTransformedUserType TransformedUserType Getter
func (r AlibabalatourstrategyshowAPIRequest) GetTransformedUserType() string {
	return r._transformedUserType
}

// SetUserNick is UserNick Setter
// 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabalatourstrategyshowAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r AlibabalatourstrategyshowAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetUserId is UserId Setter
// 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabalatourstrategyshowAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalatourstrategyshowAPIRequest) GetUserId() string {
	return r._userId
}

// SetStrategyCode is StrategyCode Setter
// 投放计划code
func (r *AlibabalatourstrategyshowAPIRequest) SetStrategyCode(_strategyCode string) error {
	r._strategyCode = _strategyCode
	r.Set("strategy_code", _strategyCode)
	return nil
}

// GetStrategyCode StrategyCode Getter
func (r AlibabalatourstrategyshowAPIRequest) GetStrategyCode() string {
	return r._strategyCode
}

// SetUserType is UserType Setter
// 当面账户类型
func (r *AlibabalatourstrategyshowAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabalatourstrategyshowAPIRequest) GetUserType() string {
	return r._userType
}

// SetOpenid is Openid Setter
// openid
func (r *AlibabalatourstrategyshowAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r AlibabalatourstrategyshowAPIRequest) GetOpenid() string {
	return r._openid
}

// SetPageSize is PageSize Setter
// 每页权益数
func (r *AlibabalatourstrategyshowAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabalatourstrategyshowAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当面分页
func (r *AlibabalatourstrategyshowAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabalatourstrategyshowAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetWithTestBenefit is WithTestBenefit Setter
// 带出测试权益
func (r *AlibabalatourstrategyshowAPIRequest) SetWithTestBenefit(_withTestBenefit bool) error {
	r._withTestBenefit = _withTestBenefit
	r.Set("with_test_benefit", _withTestBenefit)
	return nil
}

// GetWithTestBenefit WithTestBenefit Getter
func (r AlibabalatourstrategyshowAPIRequest) GetWithTestBenefit() bool {
	return r._withTestBenefit
}

// SetNeedIdentifyRisk is NeedIdentifyRisk Setter
// 是否需要调用安全校验服务
func (r *AlibabalatourstrategyshowAPIRequest) SetNeedIdentifyRisk(_needIdentifyRisk bool) error {
	r._needIdentifyRisk = _needIdentifyRisk
	r.Set("need_identify_risk", _needIdentifyRisk)
	return nil
}

// GetNeedIdentifyRisk NeedIdentifyRisk Getter
func (r AlibabalatourstrategyshowAPIRequest) GetNeedIdentifyRisk() bool {
	return r._needIdentifyRisk
}

// SetSkipWithHadWin is SkipWithHadWin Setter
// 不带出hadWin状态
func (r *AlibabalatourstrategyshowAPIRequest) SetSkipWithHadWin(_skipWithHadWin bool) error {
	r._skipWithHadWin = _skipWithHadWin
	r.Set("skip_with_had_win", _skipWithHadWin)
	return nil
}

// GetSkipWithHadWin SkipWithHadWin Getter
func (r AlibabalatourstrategyshowAPIRequest) GetSkipWithHadWin() bool {
	return r._skipWithHadWin
}

// SetFilterEmptyInventory is FilterEmptyInventory Setter
// 过滤无库存权益
func (r *AlibabalatourstrategyshowAPIRequest) SetFilterEmptyInventory(_filterEmptyInventory bool) error {
	r._filterEmptyInventory = _filterEmptyInventory
	r.Set("filter_empty_inventory", _filterEmptyInventory)
	return nil
}

// GetFilterEmptyInventory FilterEmptyInventory Getter
func (r AlibabalatourstrategyshowAPIRequest) GetFilterEmptyInventory() bool {
	return r._filterEmptyInventory
}

// SetFilterCrowd is FilterCrowd Setter
// 过滤人群
func (r *AlibabalatourstrategyshowAPIRequest) SetFilterCrowd(_filterCrowd bool) error {
	r._filterCrowd = _filterCrowd
	r.Set("filter_crowd", _filterCrowd)
	return nil
}

// GetFilterCrowd FilterCrowd Getter
func (r AlibabalatourstrategyshowAPIRequest) GetFilterCrowd() bool {
	return r._filterCrowd
}

// SetWithStrategyInstance is WithStrategyInstance Setter
// 是否需要投放计划维度的权益核销实例
func (r *AlibabalatourstrategyshowAPIRequest) SetWithStrategyInstance(_withStrategyInstance bool) error {
	r._withStrategyInstance = _withStrategyInstance
	r.Set("with_strategy_instance", _withStrategyInstance)
	return nil
}

// GetWithStrategyInstance WithStrategyInstance Getter
func (r AlibabalatourstrategyshowAPIRequest) GetWithStrategyInstance() bool {
	return r._withStrategyInstance
}

// SetWithBenefitInstance is WithBenefitInstance Setter
// 是否需要权益维度的权益核销实例
func (r *AlibabalatourstrategyshowAPIRequest) SetWithBenefitInstance(_withBenefitInstance bool) error {
	r._withBenefitInstance = _withBenefitInstance
	r.Set("with_benefit_instance", _withBenefitInstance)
	return nil
}

// GetWithBenefitInstance WithBenefitInstance Getter
func (r AlibabalatourstrategyshowAPIRequest) GetWithBenefitInstance() bool {
	return r._withBenefitInstance
}
