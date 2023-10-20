package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaocooperationgetAPIRequest 供应商或分销商获取合作关系信息 API请求
// taobao.fenxiao.cooperation.get
//
// 获取供应商的合作关系信息
type TaobaofenxiaocooperationgetAPIRequest struct {
	model.Params
	// 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
	_status string
	// 合作开始时间yyyy-MM-dd HH:mm:ss
	_startDate string
	// 合作结束时间yyyy-MM-dd HH:mm:ss
	_endDate string
	// 分销方式：AGENT(代销) 、DEALER（经销）
	_tradeType string
	// 渠道code
	_channelCode string
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
	// 1是供应商，2 是分销商
	_roleType int64
}

// NewTaobaofenxiaocooperationgetRequest 初始化TaobaofenxiaocooperationgetAPIRequest对象
func NewTaobaofenxiaocooperationgetRequest() *TaobaofenxiaocooperationgetAPIRequest {
	return &TaobaofenxiaocooperationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaocooperationgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.cooperation.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaocooperationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaocooperationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
func (r *TaobaofenxiaocooperationgetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetStatus() string {
	return r._status
}

// SetStartDate is StartDate Setter
// 合作开始时间yyyy-MM-dd HH:mm:ss
func (r *TaobaofenxiaocooperationgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 合作结束时间yyyy-MM-dd HH:mm:ss
func (r *TaobaofenxiaocooperationgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetTradeType is TradeType Setter
// 分销方式：AGENT(代销) 、DEALER（经销）
func (r *TaobaofenxiaocooperationgetAPIRequest) SetTradeType(_tradeType string) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetTradeType() string {
	return r._tradeType
}

// SetChannelCode is ChannelCode Setter
// 渠道code
func (r *TaobaofenxiaocooperationgetAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetChannelCode() string {
	return r._channelCode
}

// SetPageNo is PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaofenxiaocooperationgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaofenxiaocooperationgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetRoleType is RoleType Setter
// 1是供应商，2 是分销商
func (r *TaobaofenxiaocooperationgetAPIRequest) SetRoleType(_roleType int64) error {
	r._roleType = _roleType
	r.Set("role_type", _roleType)
	return nil
}

// GetRoleType RoleType Getter
func (r TaobaofenxiaocooperationgetAPIRequest) GetRoleType() int64 {
	return r._roleType
}
