package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupeffectGetAPIRequest 推广组效果报表数据对象 API请求
// taobao.simba.rpt.adgroupeffect.get
//
// 推广组效果报表数据对象
type TaobaoSimbaRptAdgroupeffectGetAPIRequest struct {
	model.Params
	// 权限校验参数
	_subwayToken string
	// 昵称
	_nick string
	// 推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH ）可以一次取多个例如：SEARCH,CAT
	_searchType string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoSimbaRptAdgroupeffectGetRequest 初始化TaobaoSimbaRptAdgroupeffectGetAPIRequest对象
func NewTaobaoSimbaRptAdgroupeffectGetRequest() *TaobaoSimbaRptAdgroupeffectGetAPIRequest {
	return &TaobaoSimbaRptAdgroupeffectGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.adgroupeffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// Get SubwayToken Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// Set is Nick Setter
// 昵称
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// Set is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH ）可以一次取多个例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// Get SearchType Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetSearchType() string {
	return r._searchType
}

// Set is Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetSource() string {
	return r._source
}

// Set is PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupeffectGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoSimbaRptAdgroupeffectGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
