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
type TaobaoSimbaRptCusteffectGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 开始时间，格式yyyy-mm-dd
    startTime   string
    // 结束时间，格式yyyy-mm-dd
    endTime   string
    // 权限校验参数
    subwayToken   string
    // 页码
    pageNo   int64
    // 每页大小
    pageSize   int64
    // 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    source   string
}

// 初始化TaobaoSimbaRptCusteffectGetRequest对象
func NewTaobaoSimbaRptCusteffectGetRequest() *TaobaoSimbaRptCusteffectGetRequest{
    return &TaobaoSimbaRptCusteffectGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCusteffectGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.custeffect.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCusteffectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaRptCusteffectGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptCusteffectGetRequest) GetNick() string {
    return r.nick
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCusteffectGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptCusteffectGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCusteffectGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptCusteffectGetRequest) GetEndTime() string {
    return r.endTime
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptCusteffectGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptCusteffectGetRequest) GetSubwayToken() string {
    return r.subwayToken
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptCusteffectGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptCusteffectGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCusteffectGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptCusteffectGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// Source Setter
// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptCusteffectGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptCusteffectGetRequest) GetSource() string {
    return r.source
}
