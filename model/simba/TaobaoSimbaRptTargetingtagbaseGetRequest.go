package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向基础报表 APIRequest
taobao.simba.rpt.targetingtagbase.get

获取定向基础报表
*/
type TaobaoSimbaRptTargetingtagbaseGetRequest struct {
    model.Params

    // 被操作者昵称
    nick   string 

    // 计划id
    campaignId   int64 

    // 推广组id
    adgroupId   int64 

    // 起始时间
    startTime   string 

    // 结束时间
    endTime   string 

    // 分页大小
    pageSize   int64 

    // 页码
    pageNumber   int64 

}

func NewTaobaoSimbaRptTargetingtagbaseGetRequest() *TaobaoSimbaRptTargetingtagbaseGetRequest{
    return &TaobaoSimbaRptTargetingtagbaseGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.targetingtagbase.get"
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaRptTargetingtagbaseGetRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

func (r TaobaoSimbaRptTargetingtagbaseGetRequest) GetPageNumber() int64 {
    return r.pageNumber
}

