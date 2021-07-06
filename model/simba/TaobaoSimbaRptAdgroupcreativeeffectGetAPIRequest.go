package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest 推广组下的创意报表效果数据查询(汇总数据，不分类型) API请求
// taobao.simba.rpt.adgroupcreativeeffect.get
//
// 推广组下的创意报表效果数据查询(汇总数据，不分类型)
type TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest struct {
	model.Params
	// 权限验证信息
	_subwayToken string
	// 昵称
	_nick string
	// 开始日期，格式yyyy-mm-dd
	_startTime string
	// 结束日期，格式yyyy-mm-dd
	_endTime string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
	_searchType string
	// 查询推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoSimbaRptAdgroupcreativeeffectGetRequest 初始化TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest对象
func NewTaobaoSimbaRptAdgroupcreativeeffectGetRequest() *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest {
	return &TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.adgroupcreativeeffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubwayToken is SubwayToken Setter
// 权限验证信息
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSource is Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetSource() string {
	return r._source
}

// SetSearchType is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// GetSearchType SearchType Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetSearchType() string {
	return r._searchType
}

// SetCampaignId is CampaignId Setter
// 查询推广计划id
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
