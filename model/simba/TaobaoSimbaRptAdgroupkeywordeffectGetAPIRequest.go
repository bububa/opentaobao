package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest 推广组下的词效果报表数据查询(明细数据不分类型查询) API请求
// taobao.simba.rpt.adgroupkeywordeffect.get
//
// 推广组下的词效果报表数据查询(明细数据不分类型查询)
type TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划ID
	_campaignId int64
	// 推广组ID
	_adgroupId int64
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 权限校验参数
	_subwayToken string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
	_searchType string
}

// NewTaobaoSimbaRptAdgroupkeywordeffectGetRequest 初始化TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest对象
func NewTaobaoSimbaRptAdgroupkeywordeffectGetRequest() *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest {
	return &TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.adgroupkeywordeffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 推广计划ID
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is AdgroupId Setter
// 推广组ID
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// Set is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetSource() string {
	return r._source
}

// Set is SubwayToken Setter
// 权限校验参数
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// Get SubwayToken Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// Set is PageNo Setter
// 页码
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页大小
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
func (r *TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// Get SearchType Getter
func (r TaobaoSimbaRptAdgroupkeywordeffectGetAPIRequest) GetSearchType() string {
	return r._searchType
}
