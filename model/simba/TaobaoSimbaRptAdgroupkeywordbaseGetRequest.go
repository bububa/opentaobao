package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的词基础报表数据查询(明细数据不分类型查询) API请求
taobao.simba.rpt.adgroupkeywordbase.get

推广组下的词基础报表数据查询(明细数据不分类型查询)
*/
type TaobaoSimbaRptAdgroupkeywordbaseGetRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广计划ID
    campaignId   int64
    // 推广组ID
    adgroupId   int64
    // 开始时间，格式yyyy-mm-dd
    startTime   string
    // 结束时间，格式yyyy-mm-dd
    endTime   string
    // 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    source   string
    // 权限校验参数
    subwayToken   string
    // 页码
    pageNo   int64
    // 每页大小
    pageSize   int64
    // 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
    searchType   string
}

// 初始化TaobaoSimbaRptAdgroupkeywordbaseGetRequest对象
func NewTaobaoSimbaRptAdgroupkeywordbaseGetRequest() *TaobaoSimbaRptAdgroupkeywordbaseGetRequest{
    return &TaobaoSimbaRptAdgroupkeywordbaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupkeywordbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetNick() string {
    return r.nick
}
// CampaignId Setter
// 推广计划ID
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetEndTime() string {
    return r.endTime
}
// Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetSource() string {
    return r.source
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetSubwayToken() string {
    return r.subwayToken
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetSearchType() string {
    return r.searchType
}
