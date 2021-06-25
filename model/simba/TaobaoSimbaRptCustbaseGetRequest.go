package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客户账户报表基础数据对象 APIRequest
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

func NewTaobaoSimbaRptCustbaseGetRequest() *TaobaoSimbaRptCustbaseGetRequest{
    return &TaobaoSimbaRptCustbaseGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.custbase.get"
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptCustbaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}

func (r *TaobaoSimbaRptCustbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptCustbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptCustbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptCustbaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaRptCustbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaRptCustbaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoSimbaRptCustbaseGetRequest) GetSource() string {
    return r.source
}

