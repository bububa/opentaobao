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
type TaobaoSimbaRptCustbaseGetRequest struct {
    model.Params
    // 权限验证信息
    subwayToken   string
    // 昵称
    nick   string
    // 开始日期，格式yyyy-mm-dd
    startTime   string
    // 结束日期，格式yyyy-mm-dd
    endTime   string
    // 页码
    pageNo   int64
    // 每页大小
    pageSize   int64
    // 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    source   string
}

// 初始化TaobaoSimbaRptCustbaseGetRequest对象
func NewTaobaoSimbaRptCustbaseGetRequest() *TaobaoSimbaRptCustbaseGetRequest{
    return &TaobaoSimbaRptCustbaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCustbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.custbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCustbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限验证信息
func (r *TaobaoSimbaRptCustbaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptCustbaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptCustbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptCustbaseGetRequest) GetNick() string {
    return r.nick
}
// StartTime Setter
// 开始日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCustbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptCustbaseGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCustbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptCustbaseGetRequest) GetEndTime() string {
    return r.endTime
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptCustbaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptCustbaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCustbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptCustbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// Source Setter
// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptCustbaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptCustbaseGetRequest) GetSource() string {
    return r.source
}
