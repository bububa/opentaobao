package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询客服在线状态 API请求
taobao.qianniu.cloudkefu.onlinestatuslog.get

按天查询客服账号的在线状态记录。如：登录，下线，挂起等
有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询
*/
type TaobaoQianniuCloudkefuOnlinestatuslogGetRequest struct {
    model.Params
    // 子帐号列表，最多10个
    accountIds   []int64
    // 查询开始日期，只有日期有效，时间忽略
    startDate   string
    // 查询结束日期，只有日期有效，时间忽略
    endDate   string
}

// 初始化TaobaoQianniuCloudkefuOnlinestatuslogGetRequest对象
func NewTaobaoQianniuCloudkefuOnlinestatuslogGetRequest() *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest{
    return &TaobaoQianniuCloudkefuOnlinestatuslogGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetApiMethodName() string {
    return "taobao.qianniu.cloudkefu.onlinestatuslog.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountIds Setter
// 子帐号列表，最多10个
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetAccountIds(accountIds []int64) error {
    r.accountIds = accountIds
    r.Set("account_ids", accountIds)
    return nil
}

// AccountIds Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetAccountIds() []int64 {
    return r.accountIds
}
// StartDate Setter
// 查询开始日期，只有日期有效，时间忽略
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 查询结束日期，只有日期有效，时间忽略
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetEndDate() string {
    return r.endDate
}
