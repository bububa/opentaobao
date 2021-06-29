package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的创意报表效果数据查询(汇总数据，不分类型) API请求
taobao.simba.rpt.adgroupcreativeeffect.get

推广组下的创意报表效果数据查询(汇总数据，不分类型)
*/
type TaobaoSimbaRptAdgroupcreativeeffectGetRequest struct {
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

// 初始化TaobaoSimbaRptAdgroupcreativeeffectGetRequest对象
func NewTaobaoSimbaRptAdgroupcreativeeffectGetRequest() *TaobaoSimbaRptAdgroupcreativeeffectGetRequest{
    return &TaobaoSimbaRptAdgroupcreativeeffectGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupcreativeeffect.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限验证信息
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetSubwayToken() string {
    return r.subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetNick() string {
    return r.nick
}
// StartTime Setter
// 开始日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetEndTime() string {
    return r.endTime
}
// CampaignId Setter
// 查询推广计划id
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetCampaignId() int64 {
    return r.campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
// Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetSource() string {
    return r.source
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetRequest) GetSearchType() string {
    return r.searchType
}
