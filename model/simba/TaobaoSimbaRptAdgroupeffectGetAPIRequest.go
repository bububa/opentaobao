package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptadgroupeffectgetAPIRequest 推广组效果报表数据对象 API请求
// taobao.simba.rpt.adgroupeffect.get
//
// 推广组效果报表数据对象
type TaobaosimbarptadgroupeffectgetAPIRequest struct {
	model.Params
	// 权限校验参数
	_subwayToken string
	// 昵称
	_nick string
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH ）可以一次取多个例如：SEARCH,CAT
	_searchType string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
	_source string
	// 推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaosimbarptadgroupeffectgetRequest 初始化TaobaosimbarptadgroupeffectgetAPIRequest对象
func NewTaobaosimbarptadgroupeffectgetRequest() *TaobaosimbarptadgroupeffectgetAPIRequest {
	return &TaobaosimbarptadgroupeffectgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.adgroupeffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubwayToken is SubwayToken Setter
// 权限校验参数
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSearchType is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT,&lt;br/&gt;定向投放：NOSEARCH ）可以一次取多个例如：SEARCH,CAT
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// GetSearchType SearchType Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetSearchType() string {
	return r._searchType
}

// SetSource is Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetSource() string {
	return r._source
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbarptadgroupeffectgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarptadgroupeffectgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
