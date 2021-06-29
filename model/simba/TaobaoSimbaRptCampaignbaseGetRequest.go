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
    _subwayToken   string
    // 昵称
    _nick   string
    // 开始时间，格式yyyy-mm-dd
    _startTime   string
    // 结束时间，格式yyyy-mm-dd
    _endTime   string
    // 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
    _searchType   string
    // 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
    _source   string
    // 页码
    _pageNo   int64
    // 每页大小
    _pageSize   int64
    // 推广计划id
    _campaignId   int64
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
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSubwayToken() string {
    return r._subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetEndTime() string {
    return r._endTime
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSearchType(_searchType string) error {
    r._searchType = _searchType
    r.Set("search_type", _searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSearchType() string {
    return r._searchType
}
// Source Setter
// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetSource() string {
    return r._source
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptCampaignbaseGetRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptCampaignbaseGetRequest) GetCampaignId() int64 {
    return r._campaignId
}
