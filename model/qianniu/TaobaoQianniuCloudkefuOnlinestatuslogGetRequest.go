package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询客服在线状态 APIRequest
taobao.qianniu.cloudkefu.onlinestatuslog.get

按天查询客服账号的在线状态记录。如：登录，下线，挂起等
有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询
*/
type TaobaoQianniuCloudkefuOnlinestatuslogGetRequest struct {
    model.Params

    // 子帐号列表，最多10个
    accountIds   []Number 

    // 查询开始日期，只有日期有效，时间忽略
    startDate   string 

    // 查询结束日期，只有日期有效，时间忽略
    endDate   string 

}

func NewTaobaoQianniuCloudkefuOnlinestatuslogGetRequest() *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest{
    return &TaobaoQianniuCloudkefuOnlinestatuslogGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetApiMethodName() string {
    return "taobao.qianniu.cloudkefu.onlinestatuslog.get"
}

func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetAccountIds(accountIds []Number) error {
    r.accountIds = accountIds
    r.Set("account_ids", accountIds)
    return nil
}

func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetAccountIds() []Number {
    return r.accountIds
}

func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetEndDate() string {
    return r.endDate
}

