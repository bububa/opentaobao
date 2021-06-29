package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的词效果报表数据查询(明细数据不分类型查询) API请求
taobao.simba.rpt.adgroupkeywordeffect.get

推广组下的词效果报表数据查询(明细数据不分类型查询)
*/
type TaobaoSimbaRptAdgroupkeywordeffectGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广计划ID
    _campaignId   int64
    // 推广组ID
    _adgroupId   int64
    // 开始时间，格式yyyy-mm-dd
    _startTime   string
    // 结束时间，格式yyyy-mm-dd
    _endTime   string
    // 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    _source   string
    // 权限校验参数
    _subwayToken   string
    // 页码
    _pageNo   int64
    // 每页大小
    _pageSize   int64
    // 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
    _searchType   string
}

// 初始化TaobaoSimbaRptAdgroupkeywordeffectGetRequest对象
func NewTaobaoSimbaRptAdgroupkeywordeffectGetRequest() *TaobaoSimbaRptAdgroupkeywordeffectGetRequest{
    return &TaobaoSimbaRptAdgroupkeywordeffectGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupkeywordeffect.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划ID
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetEndTime() string {
    return r._endTime
}
// Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetSource() string {
    return r._source
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetSubwayToken() string {
    return r._subwayToken
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetRequest) SetSearchType(_searchType string) error {
    r._searchType = _searchType
    r.Set("search_type", _searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetRequest) GetSearchType() string {
    return r._searchType
}
