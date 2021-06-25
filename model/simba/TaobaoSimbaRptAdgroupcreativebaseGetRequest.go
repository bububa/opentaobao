package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下创意报表基础数据查询(汇总数据，不分类型) APIRequest
taobao.simba.rpt.adgroupcreativebase.get

推广组下创意报表基础数据查询(汇总数据，不分类型)
*/
type TaobaoSimbaRptAdgroupcreativebaseGetRequest struct {
    model.Params

    // 权限验证信息
    subwayToken   string 

    // 昵称
    nick   string 

    // 开始日期，格式yyyy-mm-dd
    startTime   string 

    // 结束日期，格式yyyy-mm-dd
    endTime   string 

    // 查询推广计划id
    campaignId   int64 

    // 推广组id
    adgroupId   int64 

    // 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    source   string 

    // 页码
    pageNo   int64 

    // 每页大小
    pageSize   int64 

    // 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
    searchType   string 

}

func NewTaobaoSimbaRptAdgroupcreativebaseGetRequest() *TaobaoSimbaRptAdgroupcreativebaseGetRequest{
    return &TaobaoSimbaRptAdgroupcreativebaseGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupcreativebase.get"
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetSource() string {
    return r.source
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetSearchType() string {
    return r.searchType
}

