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
    _accountIds   []int64
    // 查询开始日期，只有日期有效，时间忽略
    _startDate   string
    // 查询结束日期，只有日期有效，时间忽略
    _endDate   string
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
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetAccountIds(_accountIds []int64) error {
    r._accountIds = _accountIds
    r.Set("account_ids", _accountIds)
    return nil
}

// AccountIds Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetAccountIds() []int64 {
    return r._accountIds
}
// StartDate Setter
// 查询开始日期，只有日期有效，时间忽略
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 查询结束日期，只有日期有效，时间忽略
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetRequest) GetEndDate() string {
    return r._endDate
}
