package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组基础报表数据对象 APIRequest
taobao.simba.rpt.adgroupbase.get

推广组基础报表数据对象
*/
type TaobaoSimbaRptAdgroupbaseGetRequest struct {
    model.Params

    // 权限校验参数
    subwayToken   string 

    // 昵称
    nick   string 

    // 推广计划id
    campaignId   int64 

    // 推广组id
    adgroupId   int64 

    // 开始时间，格式yyyy-mm-dd
    startTime   string 

    // 结束时间，格式yyyy-mm-dd
    endTime   string 

    // 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH）可以一次取多个例如：SEARCH,CAT
    searchType   string 

    // 页码
    pageNo   int64 

    // 每页大小
    pageSize   int64 

    // 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
    source   string 

}

func NewTaobaoSimbaRptAdgroupbaseGetRequest() *TaobaoSimbaRptAdgroupbaseGetRequest{
    return &TaobaoSimbaRptAdgroupbaseGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupbase.get"
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetSearchType() string {
    return r.searchType
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaRptAdgroupbaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoSimbaRptAdgroupbaseGetRequest) GetSource() string {
    return r.source
}

