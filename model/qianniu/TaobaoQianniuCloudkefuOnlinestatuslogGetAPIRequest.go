package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniucloudkefuonlinestatusloggetAPIRequest 查询客服在线状态 API请求
// taobao.qianniu.cloudkefu.onlinestatuslog.get
//
// 按天查询客服账号的在线状态记录。如：登录，下线，挂起等
// 有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询
type TaobaoqianniucloudkefuonlinestatusloggetAPIRequest struct {
	model.Params
	// 子帐号列表，最多10个
	_accountIds []string
	// 查询开始日期，只有日期有效，时间忽略
	_startDate string
	// 查询结束日期，只有日期有效，时间忽略
	_endDate string
}

// NewTaobaoqianniucloudkefuonlinestatusloggetRequest 初始化TaobaoqianniucloudkefuonlinestatusloggetAPIRequest对象
func NewTaobaoqianniucloudkefuonlinestatusloggetRequest() *TaobaoqianniucloudkefuonlinestatusloggetAPIRequest {
	return &TaobaoqianniucloudkefuonlinestatusloggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.cloudkefu.onlinestatuslog.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountIds is AccountIds Setter
// 子帐号列表，最多10个
func (r *TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) SetAccountIds(_accountIds []string) error {
	r._accountIds = _accountIds
	r.Set("account_ids", _accountIds)
	return nil
}

// GetAccountIds AccountIds Getter
func (r TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) GetAccountIds() []string {
	return r._accountIds
}

// SetStartDate is StartDate Setter
// 查询开始日期，只有日期有效，时间忽略
func (r *TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 查询结束日期，只有日期有效，时间忽略
func (r *TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoqianniucloudkefuonlinestatusloggetAPIRequest) GetEndDate() string {
	return r._endDate
}
