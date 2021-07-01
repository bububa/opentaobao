package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户账户报表基础数据对象 API请求
taobao.simba.rpt.custbase.get

客户账户报表基础数据对象
*/
type TaobaoSimbaRptCustbaseGetAPIRequest struct {
    model.Params
    // 权限验证信息
    _subwayToken   string
    // 昵称
    _nick   string
    // 开始日期，格式yyyy-mm-dd
    _startTime   string
    // 结束日期，格式yyyy-mm-dd
    _endTime   string
    // 页码
    _pageNo   int64
    // 每页大小
    _pageSize   int64
    // 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    _source   string
}

// 初始化TaobaoSimbaRptCustbaseGetAPIRequest对象
func NewTaobaoSimbaRptCustbaseGetRequest() *TaobaoSimbaRptCustbaseGetAPIRequest{
    return &TaobaoSimbaRptCustbaseGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.custbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限验证信息
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetSubwayToken() string {
    return r._subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 开始日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetEndTime() string {
    return r._endTime
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Source Setter
// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptCustbaseGetAPIRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptCustbaseGetAPIRequest) GetSource() string {
    return r._source
}
