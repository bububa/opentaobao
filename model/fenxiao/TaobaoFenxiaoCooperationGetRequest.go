package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商或分销商获取合作关系信息 API请求
taobao.fenxiao.cooperation.get

获取供应商的合作关系信息
*/
type TaobaoFenxiaoCooperationGetRequest struct {
    model.Params
    // 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
    _status   string
    // 合作开始时间yyyy-MM-dd HH:mm:ss
    _startDate   string
    // 合作结束时间yyyy-MM-dd HH:mm:ss
    _endDate   string
    // 分销方式：AGENT(代销) 、DEALER（经销）
    _tradeType   string
    // 页码（大于0的整数，默认1）
    _pageNo   int64
    // 每页记录数（默认20，最大50）
    _pageSize   int64
    // 渠道code
    _channelCode   string
    // 1是供应商，2 是分销商
    _roleType   int64
}

// 初始化TaobaoFenxiaoCooperationGetRequest对象
func NewTaobaoFenxiaoCooperationGetRequest() *TaobaoFenxiaoCooperationGetRequest{
    return &TaobaoFenxiaoCooperationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoCooperationGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.cooperation.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoCooperationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
func (r *TaobaoFenxiaoCooperationGetRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetStatus() string {
    return r._status
}
// StartDate Setter
// 合作开始时间yyyy-MM-dd HH:mm:ss
func (r *TaobaoFenxiaoCooperationGetRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 合作结束时间yyyy-MM-dd HH:mm:ss
func (r *TaobaoFenxiaoCooperationGetRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetEndDate() string {
    return r._endDate
}
// TradeType Setter
// 分销方式：AGENT(代销) 、DEALER（经销）
func (r *TaobaoFenxiaoCooperationGetRequest) SetTradeType(_tradeType string) error {
    r._tradeType = _tradeType
    r.Set("trade_type", _tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetTradeType() string {
    return r._tradeType
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoCooperationGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoCooperationGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// ChannelCode Setter
// 渠道code
func (r *TaobaoFenxiaoCooperationGetRequest) SetChannelCode(_channelCode string) error {
    r._channelCode = _channelCode
    r.Set("channel_code", _channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetChannelCode() string {
    return r._channelCode
}
// RoleType Setter
// 1是供应商，2 是分销商
func (r *TaobaoFenxiaoCooperationGetRequest) SetRoleType(_roleType int64) error {
    r._roleType = _roleType
    r.Set("role_type", _roleType)
    return nil
}

// RoleType Getter
func (r TaobaoFenxiaoCooperationGetRequest) GetRoleType() int64 {
    return r._roleType
}
