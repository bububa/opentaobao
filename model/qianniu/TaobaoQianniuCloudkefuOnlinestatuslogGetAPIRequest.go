package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest 查询客服在线状态 API请求
// taobao.qianniu.cloudkefu.onlinestatuslog.get
//
// 按天查询客服账号的在线状态记录。如：登录，下线，挂起等
// 有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询
type TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest struct {
	model.Params
	// 子帐号列表，最多10个
	_accountIds []string
	// 查询开始日期，只有日期有效，时间忽略
	_startDate string
	// 查询结束日期，只有日期有效，时间忽略
	_endDate string
}

// NewTaobaoQianniuCloudkefuOnlinestatuslogGetRequest 初始化TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest对象
func NewTaobaoQianniuCloudkefuOnlinestatuslogGetRequest() *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest {
	return &TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.cloudkefu.onlinestatuslog.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAccountIds is AccountIds Setter
// 子帐号列表，最多10个
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) SetAccountIds(_accountIds []string) error {
	r._accountIds = _accountIds
	r.Set("account_ids", _accountIds)
	return nil
}

// GetAccountIds AccountIds Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) GetAccountIds() []string {
	return r._accountIds
}

// SetStartDate is StartDate Setter
// 查询开始日期，只有日期有效，时间忽略
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询结束日期，只有日期有效，时间忽略
func (r *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoQianniuCloudkefuOnlinestatuslogGetAPIRequest) GetEndDate() string {
	return r._endDate
}
