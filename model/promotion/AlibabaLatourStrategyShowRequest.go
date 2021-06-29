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
    withTestBenefit   bool
    // 渠道
    channel   string
    // 每页权益数
    pageSize   int64
    // 要转换的账户类型
    transformedUserType   string
    // 是否需要调用安全校验服务
    needIdentifyRisk   bool
    // 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    userNick   string
    // 不带出hadWin状态
    skipWithHadWin   bool
    // 过滤无库存权益
    filterEmptyInventory   bool
    // 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    userId   string
    // 投放计划code
    strategyCode   string
    // 当面账户类型
    userType   string
    // 当面分页
    currentPage   int64
    // 过滤人群
    filterCrowd   bool
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
func (r *AlibabaLatourStrategyShowRequest) SetWithTestBenefit(withTestBenefit bool) error {
    r.withTestBenefit = withTestBenefit
    r.Set("with_test_benefit", withTestBenefit)
    return nil
}

// WithTestBenefit Getter
func (r AlibabaLatourStrategyShowRequest) GetWithTestBenefit() bool {
    return r.withTestBenefit
}
// Channel Setter
// 渠道
func (r *AlibabaLatourStrategyShowRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r AlibabaLatourStrategyShowRequest) GetChannel() string {
    return r.channel
}
// PageSize Setter
// 每页权益数
func (r *AlibabaLatourStrategyShowRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaLatourStrategyShowRequest) GetPageSize() int64 {
    return r.pageSize
}
// TransformedUserType Setter
// 要转换的账户类型
func (r *AlibabaLatourStrategyShowRequest) SetTransformedUserType(transformedUserType string) error {
    r.transformedUserType = transformedUserType
    r.Set("transformed_user_type", transformedUserType)
    return nil
}

// TransformedUserType Getter
func (r AlibabaLatourStrategyShowRequest) GetTransformedUserType() string {
    return r.transformedUserType
}
// NeedIdentifyRisk Setter
// 是否需要调用安全校验服务
func (r *AlibabaLatourStrategyShowRequest) SetNeedIdentifyRisk(needIdentifyRisk bool) error {
    r.needIdentifyRisk = needIdentifyRisk
    r.Set("need_identify_risk", needIdentifyRisk)
    return nil
}

// NeedIdentifyRisk Getter
func (r AlibabaLatourStrategyShowRequest) GetNeedIdentifyRisk() bool {
    return r.needIdentifyRisk
}
// UserNick Setter
// 用户昵称，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyShowRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r AlibabaLatourStrategyShowRequest) GetUserNick() string {
    return r.userNick
}
// SkipWithHadWin Setter
// 不带出hadWin状态
func (r *AlibabaLatourStrategyShowRequest) SetSkipWithHadWin(skipWithHadWin bool) error {
    r.skipWithHadWin = skipWithHadWin
    r.Set("skip_with_had_win", skipWithHadWin)
    return nil
}

// SkipWithHadWin Getter
func (r AlibabaLatourStrategyShowRequest) GetSkipWithHadWin() bool {
    return r.skipWithHadWin
}
// FilterEmptyInventory Setter
// 过滤无库存权益
func (r *AlibabaLatourStrategyShowRequest) SetFilterEmptyInventory(filterEmptyInventory bool) error {
    r.filterEmptyInventory = filterEmptyInventory
    r.Set("filter_empty_inventory", filterEmptyInventory)
    return nil
}

// FilterEmptyInventory Getter
func (r AlibabaLatourStrategyShowRequest) GetFilterEmptyInventory() bool {
    return r.filterEmptyInventory
}
// UserId Setter
// 用户id，除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyShowRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaLatourStrategyShowRequest) GetUserId() string {
    return r.userId
}
// StrategyCode Setter
// 投放计划code
func (r *AlibabaLatourStrategyShowRequest) SetStrategyCode(strategyCode string) error {
    r.strategyCode = strategyCode
    r.Set("strategy_code", strategyCode)
    return nil
}

// StrategyCode Getter
func (r AlibabaLatourStrategyShowRequest) GetStrategyCode() string {
    return r.strategyCode
}
// UserType Setter
// 当面账户类型
func (r *AlibabaLatourStrategyShowRequest) SetUserType(userType string) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

// UserType Getter
func (r AlibabaLatourStrategyShowRequest) GetUserType() string {
    return r.userType
}
// CurrentPage Setter
// 当面分页
func (r *AlibabaLatourStrategyShowRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaLatourStrategyShowRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// FilterCrowd Setter
// 过滤人群
func (r *AlibabaLatourStrategyShowRequest) SetFilterCrowd(filterCrowd bool) error {
    r.filterCrowd = filterCrowd
    r.Set("filter_crowd", filterCrowd)
    return nil
}

// FilterCrowd Getter
func (r AlibabaLatourStrategyShowRequest) GetFilterCrowd() bool {
    return r.filterCrowd
}
