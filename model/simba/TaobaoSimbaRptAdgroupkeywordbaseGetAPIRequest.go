package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptadgroupkeywordbasegetAPIRequest 推广组下的词基础报表数据查询(明细数据不分类型查询) API请求
// taobao.simba.rpt.adgroupkeywordbase.get
//
// 推广组下的词基础报表数据查询(明细数据不分类型查询)
type TaobaosimbarptadgroupkeywordbasegetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 权限校验参数
	_subwayToken string
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
	_searchType string
	// 推广计划ID
	_campaignId int64
	// 推广组ID
	_adgroupId int64
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaosimbarptadgroupkeywordbasegetRequest 初始化TaobaosimbarptadgroupkeywordbasegetAPIRequest对象
func NewTaobaosimbarptadgroupkeywordbasegetRequest() *TaobaosimbarptadgroupkeywordbasegetAPIRequest {
	return &TaobaosimbarptadgroupkeywordbasegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.adgroupkeywordbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSource is Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetSource() string {
	return r._source
}

// SetSubwayToken is SubwayToken Setter
// 权限校验参数
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetSearchType is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// GetSearchType SearchType Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetSearchType() string {
	return r._searchType
}

// SetCampaignId is CampaignId Setter
// 推广计划ID
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组ID
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbarptadgroupkeywordbasegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarptadgroupkeywordbasegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
