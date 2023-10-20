package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptcampaignbasegetAPIRequest 推广计划报表基础数据对象 API请求
// taobao.simba.rpt.campaignbase.get
//
// 推广计划报表基础数据对象
type TaobaosimbarptcampaignbasegetAPIRequest struct {
	model.Params
	// 权限校验参数
	_subwayToken string
	// 昵称
	_nick string
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
	_searchType string
	// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 推广计划id
	_campaignId int64
}

// NewTaobaosimbarptcampaignbasegetRequest 初始化TaobaosimbarptcampaignbasegetAPIRequest对象
func NewTaobaosimbarptcampaignbasegetRequest() *TaobaosimbarptcampaignbasegetAPIRequest {
	return &TaobaosimbarptcampaignbasegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.campaignbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubwayToken is SubwayToken Setter
// 权限校验参数
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间，格式yyyy-mm-dd
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，格式yyyy-mm-dd
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSearchType is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// GetSearchType SearchType Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetSearchType() string {
	return r._searchType
}

// SetSource is Source Setter
// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetSource() string {
	return r._source
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaosimbarptcampaignbasegetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarptcampaignbasegetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
