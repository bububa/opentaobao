package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户账户报表效果数据查询（只有汇总数据，无分类数据） API请求
taobao.simba.rpt.custeffect.get

用户账户报表效果数据查询（只有汇总数据，无分类数据）
*/
type TaobaoSimbaRptCusteffectGetAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 开始时间，格式yyyy-mm-dd
    _startTime   string
    // 结束时间，格式yyyy-mm-dd
    _endTime   string
    // 权限校验参数
    _subwayToken   string
    // 页码
    _pageNo   int64
    // 每页大小
    _pageSize   int64
    // 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    _source   string
}

// 初始化TaobaoSimbaRptCusteffectGetAPIRequest对象
func NewTaobaoSimbaRptCusteffectGetRequest() *TaobaoSimbaRptCusteffectGetAPIRequest{
    return &TaobaoSimbaRptCusteffectGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.custeffect.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaRptCusteffectGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCusteffectGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCusteffectGetAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetEndTime() string {
    return r._endTime
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptCusteffectGetAPIRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetSubwayToken() string {
    return r._subwayToken
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptCusteffectGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCusteffectGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Source Setter
// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptCusteffectGetAPIRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptCusteffectGetAPIRequest) GetSource() string {
    return r._source
}
