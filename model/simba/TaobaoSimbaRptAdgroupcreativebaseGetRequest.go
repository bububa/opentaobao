package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下创意报表基础数据查询(汇总数据，不分类型) API请求
taobao.simba.rpt.adgroupcreativebase.get

推广组下创意报表基础数据查询(汇总数据，不分类型)
*/
type TaobaoSimbaRptAdgroupcreativebaseGetRequest struct {
    model.Params
    // 权限验证信息
    _subwayToken   string
    // 昵称
    _nick   string
    // 开始日期，格式yyyy-mm-dd
    _startTime   string
    // 结束日期，格式yyyy-mm-dd
    _endTime   string
    // 查询推广计划id
    _campaignId   int64
    // 推广组id
    _adgroupId   int64
    // 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
    _source   string
    // 页码
    _pageNo   int64
    // 每页大小
    _pageSize   int64
    // 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
    _searchType   string
}

// 初始化TaobaoSimbaRptAdgroupcreativebaseGetRequest对象
func NewTaobaoSimbaRptAdgroupcreativebaseGetRequest() *TaobaoSimbaRptAdgroupcreativebaseGetRequest{
    return &TaobaoSimbaRptAdgroupcreativebaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.adgroupcreativebase.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限验证信息
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetSubwayToken() string {
    return r._subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 开始日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetEndTime() string {
    return r._endTime
}
// CampaignId Setter
// 查询推广计划id
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetSource() string {
    return r._source
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupcreativebaseGetRequest) SetSearchType(_searchType string) error {
    r._searchType = _searchType
    r.Set("search_type", _searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptAdgroupcreativebaseGetRequest) GetSearchType() string {
    return r._searchType
}
