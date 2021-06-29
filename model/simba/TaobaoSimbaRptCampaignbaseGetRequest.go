package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划报表基础数据对象 API请求
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

// 初始化TaobaoSimbaRptCampaignbaseGetRequest对象
func NewTaobaoSimbaRptCampaignbaseGetRequest() *TaobaoSimbaRptCampaignbaseGetRequest{
    return &TaobaoSimbaRptCampaignbaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.campaignbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetNick() string {
    return r.nick
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetEndTime() string {
    return r.endTime
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSearchType() string {
    return r.searchType
}
// Source Setter
// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSource() string {
    return r.source
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
