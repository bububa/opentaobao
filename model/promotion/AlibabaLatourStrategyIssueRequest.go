package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益发放接口 APIRequest
alibaba.latour.strategy.issue

阿里巴巴权益平台权益发放接口
*/
type AlibabaLatourStrategyIssueRequest struct {
    model.Params

    // 扩展参数
    extraData   string 

    // 算法容灾
    failoverAlgorithmResult   bool 

    // 幂等id
    idempotentId   string 

    // 发放渠道
    channel   string 

    // 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    userId   string 

    // 转换用户类型
    transformedUserType   string 

    // 是否需要过安全
    needIdentifyRisk   bool 

    // 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    userNick   string 

    // 投放计划code
    strategyCode   string 

    // 用户类型
    userType   string 

    // 指定发放权益code
    selectedBenefitCode   string 

}

func NewAlibabaLatourStrategyIssueRequest() *AlibabaLatourStrategyIssueRequest{
    return &AlibabaLatourStrategyIssueRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLatourStrategyIssueRequest) GetApiMethodName() string {
    return "alibaba.latour.strategy.issue"
}

func (r AlibabaLatourStrategyIssueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLatourStrategyIssueRequest) SetExtraData(extraData string) error {
    r.extraData = extraData
    r.Set("extra_data", extraData)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetExtraData() string {
    return r.extraData
}

func (r *AlibabaLatourStrategyIssueRequest) SetFailoverAlgorithmResult(failoverAlgorithmResult bool) error {
    r.failoverAlgorithmResult = failoverAlgorithmResult
    r.Set("failover_algorithm_result", failoverAlgorithmResult)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetFailoverAlgorithmResult() bool {
    return r.failoverAlgorithmResult
}

func (r *AlibabaLatourStrategyIssueRequest) SetIdempotentId(idempotentId string) error {
    r.idempotentId = idempotentId
    r.Set("idempotent_id", idempotentId)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetIdempotentId() string {
    return r.idempotentId
}

func (r *AlibabaLatourStrategyIssueRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetChannel() string {
    return r.channel
}

func (r *AlibabaLatourStrategyIssueRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaLatourStrategyIssueRequest) SetTransformedUserType(transformedUserType string) error {
    r.transformedUserType = transformedUserType
    r.Set("transformed_user_type", transformedUserType)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetTransformedUserType() string {
    return r.transformedUserType
}

func (r *AlibabaLatourStrategyIssueRequest) SetNeedIdentifyRisk(needIdentifyRisk bool) error {
    r.needIdentifyRisk = needIdentifyRisk
    r.Set("need_identify_risk", needIdentifyRisk)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetNeedIdentifyRisk() bool {
    return r.needIdentifyRisk
}

func (r *AlibabaLatourStrategyIssueRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetUserNick() string {
    return r.userNick
}

func (r *AlibabaLatourStrategyIssueRequest) SetStrategyCode(strategyCode string) error {
    r.strategyCode = strategyCode
    r.Set("strategy_code", strategyCode)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetStrategyCode() string {
    return r.strategyCode
}

func (r *AlibabaLatourStrategyIssueRequest) SetUserType(userType string) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetUserType() string {
    return r.userType
}

func (r *AlibabaLatourStrategyIssueRequest) SetSelectedBenefitCode(selectedBenefitCode string) error {
    r.selectedBenefitCode = selectedBenefitCode
    r.Set("selected_benefit_code", selectedBenefitCode)
    return nil
}

func (r AlibabaLatourStrategyIssueRequest) GetSelectedBenefitCode() string {
    return r.selectedBenefitCode
}

