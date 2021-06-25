package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的词效果报表数据查询(明细数据不分类型查询) APIRequest
taobao.simba.rpt.adgroupkeywordeffect.get

推广组下的词效果报表数据查询(明细数据不分类型查询)
*/
type TaobaoSimbaRptAdgroupkeywordeffectGetRequest struct {
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

func NewTaobaoSimbaRptAdgroupkeywordeffectGetRequest() *TaobaoSimbaRptAdgroupkeywordeffectGetRequest{
    return &TaobaoSimbaRptAdgroupkeywordeffectGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupkeywordeffect.get"
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetSource() string {
    return r.source
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetSubwayToken(subwayToken string) error {
    r.subwayToken = subwayToken
    r.Set("subway_token", subwayToken)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetSubwayToken() string {
    return r.subwayToken
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetSearchType(searchType string) error {
    r.searchType = searchType
    r.Set("search_type", searchType)
    return nil
}

func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetSearchType() string {
    return r.searchType
}

