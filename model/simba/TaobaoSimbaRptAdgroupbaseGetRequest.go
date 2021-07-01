package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组基础报表数据对象 API请求
taobao.simba.rpt.adgroupbase.get

推广组基础报表数据对象
*/
type TaobaoSimbaRptAdgroupbaseGetAPIRequest struct {
    model.Params
    // 权限校验参数
    _subwayToken   string
    // 昵称
    _nick   string
    // 推广计划id
    _campaignId   int64
    // 推广组id
    _adgroupId   int64
    // 开始时间，格式yyyy-mm-dd
    _startTime   string
    // 结束时间，格式yyyy-mm-dd
    _endTime   string
    // 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH）可以一次取多个例如：SEARCH,CAT
    _searchType   string
    // 页码
    _pageNo   int64
    // 每页大小
    _pageSize   int64
    // 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
    _source   string
}

// 初始化TaobaoSimbaRptAdgroupbaseGetAPIRequest对象
func NewTaobaoSimbaRptAdgroupbaseGetRequest() *TaobaoSimbaRptAdgroupbaseGetAPIRequest{
    return &TaobaoSimbaRptAdgroupbaseGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupbase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetSubwayToken() string {
    return r._subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetEndTime() string {
    return r._endTime
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH）可以一次取多个例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetSearchType(_searchType string) error {
    r._searchType = _searchType
    r.Set("search_type", _searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetSearchType() string {
    return r._searchType
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
func (r *TaobaoSimbaRptAdgroupbaseGetAPIRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptAdgroupbaseGetAPIRequest) GetSource() string {
    return r._source
}
