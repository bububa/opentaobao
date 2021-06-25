package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划报表基础数据对象 APIRequest
taobao.simba.rpt.campaignbase.get

推广计划报表基础数据对象
*/
type TaobaoSimbaRptCampaignbaseGetRequest struct {
    model.Params

    // 权限校验参数
    subwayToken   string 

    // 昵称
    nick   string 

    // 开始时间，格式yyyy-mm-dd
    startTime   string 

    // 结束时间，格式yyyy-mm-dd
    endTime   string 

    // 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
    searchType   string 

    // 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
    source   string 

    // 页码
    pageNo   int64 

    // 每页大小
    pageSize   int64 

    // 推广计划id
    campaignId   int64 

}

func NewTaobaoSimbaRptCampaignbaseGetRequest() *TaobaoSimbaRptCampaignbaseGetRequest{
    return &TaobaoSimbaRptCampaignbaseGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.campaignbase.get"
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSearchType() string {
    return r.searchType
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSource() string {
    return r.source
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRptCampaignbaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

