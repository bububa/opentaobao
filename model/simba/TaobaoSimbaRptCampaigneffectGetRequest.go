package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划效果报表数据对象 APIRequest
taobao.simba.rpt.campaigneffect.get

推广计划效果报表数据对象
*/
type TaobaoSimbaRptCampaigneffectGetRequest struct {
    model.Params

    // 权限校验参数
    subwayToken   string 

    // 昵称
    nick   string 

    // 开始时间，格式yyyy-mm-dd
    startTime   string 

    // 结束时间，格式yyyy-mm-dd
    endTime   string 

    // 推广计划id
    campaignId   int64 

    // 报表类型（搜索：SEARCH,类目出价：CAT,定向投放：NOSEARCH 全部：ALL）
    searchType   string 

    // 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
    source   string 

    // 页码
    pageNo   int64 

    // 每页大小
    pageSize   int64 

}

func NewTaobaoSimbaRptCampaigneffectGetRequest() *TaobaoSimbaRptCampaigneffectGetRequest{
    return &TaobaoSimbaRptCampaigneffectGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.campaigneffect.get"
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetSubwayToken() string {
    return r.subwayToken
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetSearchType() string {
    return r.searchType
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetSource() string {
    return r.source
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaRptCampaigneffectGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRptCampaigneffectGetRequest) GetPageSize() int64 {
    return r.pageSize
}

