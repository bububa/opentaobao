package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划效果报表数据对象 API请求
taobao.simba.rpt.campaigneffect.get

推广计划效果报表数据对象
*/
type TaobaoSimbaRptCampaigneffectGetAPIRequest struct {
    model.Params
    // 权限校验参数
    _subwayToken   string
    // 昵称
    _nick   string
    // 开始时间，格式yyyy-mm-dd
    _startTime   string
    // 结束时间，格式yyyy-mm-dd
    _endTime   string
    // 推广计划id
    _campaignId   int64
    // 报表类型（搜索：SEARCH,类目出价：CAT,定向投放：NOSEARCH 全部：ALL）
    _searchType   string
    // 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
    _source   string
    // 页码
    _pageNo   int64
    // 每页大小
    _pageSize   int64
}

// 初始化TaobaoSimbaRptCampaigneffectGetAPIRequest对象
func NewTaobaoSimbaRptCampaigneffectGetRequest() *TaobaoSimbaRptCampaigneffectGetAPIRequest{
    return &TaobaoSimbaRptCampaigneffectGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rpt.campaigneffect.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetSubwayToken(_subwayToken string) error {
    r._subwayToken = _subwayToken
    r.Set("subway_token", _subwayToken)
    return nil
}

// SubwayToken Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetSubwayToken() string {
    return r._subwayToken
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetEndTime() string {
    return r._endTime
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT,定向投放：NOSEARCH 全部：ALL）
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetSearchType(_searchType string) error {
    r._searchType = _searchType
    r.Set("search_type", _searchType)
    return nil
}

// SearchType Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetSearchType() string {
    return r._searchType
}
// Source Setter
// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetSource() string {
    return r._source
}
// PageNo Setter
// 页码
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptCampaigneffectGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaRptCampaigneffectGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
