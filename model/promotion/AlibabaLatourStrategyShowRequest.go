package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益投放接口 API请求
alibaba.latour.strategy.show

阿里巴巴权益平台权益投放接口
*/
type AlibabaLatourStrategyShowRequest struct {
    model.Params
    // 带出测试权益
    _withTestBenefit   bool
    // 渠道
    _channel   string
    // 每页权益数
    _pageSize   int64
    // 要转换的账户类型
    _transformedUserType   string
    // 是否需要调用安全校验服务
    _needIdentifyRisk   bool
    // 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    _userNick   string
    // 不带出hadWin状态
    _skipWithHadWin   bool
    // 过滤无库存权益
    _filterEmptyInventory   bool
    // 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    _userId   string
    // 投放计划code
    _strategyCode   string
    // 当面账户类型
    _userType   string
    // 当面分页
    _currentPage   int64
    // 过滤人群
    _filterCrowd   bool
}

// 初始化AlibabaLatourStrategyShowRequest对象
func NewAlibabaLatourStrategyShowRequest() *AlibabaLatourStrategyShowRequest{
    return &AlibabaLatourStrategyShowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLatourStrategyShowRequest) GetApiMethodName() string {
    return "alibaba.latour.strategy.show"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLatourStrategyShowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WithTestBenefit Setter
// 带出测试权益
func (r *AlibabaLatourStrategyShowRequest) SetWithTestBenefit(_withTestBenefit bool) error {
    r._withTestBenefit = _withTestBenefit
    r.Set("with_test_benefit", _withTestBenefit)
    return nil
}

// WithTestBenefit Getter
func (r AlibabaLatourStrategyShowRequest) GetWithTestBenefit() bool {
    return r._withTestBenefit
}
// Channel Setter
// 渠道
func (r *AlibabaLatourStrategyShowRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r AlibabaLatourStrategyShowRequest) GetChannel() string {
    return r._channel
}
// PageSize Setter
// 每页权益数
func (r *AlibabaLatourStrategyShowRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaLatourStrategyShowRequest) GetPageSize() int64 {
    return r._pageSize
}
// TransformedUserType Setter
// 要转换的账户类型
func (r *AlibabaLatourStrategyShowRequest) SetTransformedUserType(_transformedUserType string) error {
    r._transformedUserType = _transformedUserType
    r.Set("transformed_user_type", _transformedUserType)
    return nil
}

// TransformedUserType Getter
func (r AlibabaLatourStrategyShowRequest) GetTransformedUserType() string {
    return r._transformedUserType
}
// NeedIdentifyRisk Setter
// 是否需要调用安全校验服务
func (r *AlibabaLatourStrategyShowRequest) SetNeedIdentifyRisk(_needIdentifyRisk bool) error {
    r._needIdentifyRisk = _needIdentifyRisk
    r.Set("need_identify_risk", _needIdentifyRisk)
    return nil
}

// NeedIdentifyRisk Getter
func (r AlibabaLatourStrategyShowRequest) GetNeedIdentifyRisk() bool {
    return r._needIdentifyRisk
}
// UserNick Setter
// 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyShowRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r AlibabaLatourStrategyShowRequest) GetUserNick() string {
    return r._userNick
}
// SkipWithHadWin Setter
// 不带出hadWin状态
func (r *AlibabaLatourStrategyShowRequest) SetSkipWithHadWin(_skipWithHadWin bool) error {
    r._skipWithHadWin = _skipWithHadWin
    r.Set("skip_with_had_win", _skipWithHadWin)
    return nil
}

// SkipWithHadWin Getter
func (r AlibabaLatourStrategyShowRequest) GetSkipWithHadWin() bool {
    return r._skipWithHadWin
}
// FilterEmptyInventory Setter
// 过滤无库存权益
func (r *AlibabaLatourStrategyShowRequest) SetFilterEmptyInventory(_filterEmptyInventory bool) error {
    r._filterEmptyInventory = _filterEmptyInventory
    r.Set("filter_empty_inventory", _filterEmptyInventory)
    return nil
}

// FilterEmptyInventory Getter
func (r AlibabaLatourStrategyShowRequest) GetFilterEmptyInventory() bool {
    return r._filterEmptyInventory
}
// UserId Setter
// 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyShowRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaLatourStrategyShowRequest) GetUserId() string {
    return r._userId
}
// StrategyCode Setter
// 投放计划code
func (r *AlibabaLatourStrategyShowRequest) SetStrategyCode(_strategyCode string) error {
    r._strategyCode = _strategyCode
    r.Set("strategy_code", _strategyCode)
    return nil
}

// StrategyCode Getter
func (r AlibabaLatourStrategyShowRequest) GetStrategyCode() string {
    return r._strategyCode
}
// UserType Setter
// 当面账户类型
func (r *AlibabaLatourStrategyShowRequest) SetUserType(_userType string) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r AlibabaLatourStrategyShowRequest) GetUserType() string {
    return r._userType
}
// CurrentPage Setter
// 当面分页
func (r *AlibabaLatourStrategyShowRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaLatourStrategyShowRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// FilterCrowd Setter
// 过滤人群
func (r *AlibabaLatourStrategyShowRequest) SetFilterCrowd(_filterCrowd bool) error {
    r._filterCrowd = _filterCrowd
    r.Set("filter_crowd", _filterCrowd)
    return nil
}

// FilterCrowd Getter
func (r AlibabaLatourStrategyShowRequest) GetFilterCrowd() bool {
    return r._filterCrowd
}
