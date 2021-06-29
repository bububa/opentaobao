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
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划ID
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetEndTime() string {
    return r._endTime
}
// Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetSource() string {
    return r._source
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetSubwayToken() string {
    return r._subwayToken
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupkeywordbaseGetRequest) SetSearchType(_searchType string) error {
    r._searchType = _searchType
    r.Set("search_type", _searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptAdgroupkeywordbaseGetRequest) GetSearchType() string {
    return r._searchType
}
