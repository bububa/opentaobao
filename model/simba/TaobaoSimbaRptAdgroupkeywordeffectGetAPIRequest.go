package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptadgroupkeywordeffectgetAPIRequest 推广组下的词效果报表数据查询(明细数据不分类型查询) API请求
// taobao.simba.rpt.adgroupkeywordeffect.get
//
// 推广组下的词效果报表数据查询(明细数据不分类型查询)
type TaobaosimbarptadgroupkeywordeffectgetAPIRequest struct {
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

// NewTaobaosimbarptadgroupkeywordeffectgetRequest 初始化TaobaosimbarptadgroupkeywordeffectgetAPIRequest对象
func NewTaobaosimbarptadgroupkeywordeffectgetRequest() *TaobaosimbarptadgroupkeywordeffectgetAPIRequest {
	return &TaobaosimbarptadgroupkeywordeffectgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.adgroupkeywordeffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSource is Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetSource() string {
	return r._source
}

// SetSubwayToken is SubwayToken Setter
// 权限校验参数
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetSearchType is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// GetSearchType SearchType Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetSearchType() string {
	return r._searchType
}

// SetCampaignId is CampaignId Setter
// 推广计划ID
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组ID
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbarptadgroupkeywordeffectgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarptadgroupkeywordeffectgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
