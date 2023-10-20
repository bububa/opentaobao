package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarptcampadgroupbasegetAPIRequest 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型) API请求
// taobao.simba.rpt.campadgroupbase.get
//
// 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型)
type TaobaosimbarptcampadgroupbasegetAPIRequest struct {
	model.Params
	// 权限验证信息
	_subwayToken string
	// 昵称
	_nick string
	// 开始日期,格式yyyy-mm-dd
	_startTime string
	// 结束日期,格式yyyy-mm-dd
	_endTime string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5, 汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
	_searchType string
	// 查询推广计划id
	_campaignId int64
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaosimbarptcampadgroupbasegetRequest 初始化TaobaosimbarptcampadgroupbasegetAPIRequest对象
func NewTaobaosimbarptcampadgroupbasegetRequest() *TaobaosimbarptcampadgroupbasegetAPIRequest {
	return &TaobaosimbarptcampadgroupbasegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.campadgroupbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubwayToken is SubwayToken Setter
// 权限验证信息
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetSubwayToken(_subwayToken string) error {
	r._subwayToken = _subwayToken
	r.Set("subway_token", _subwayToken)
	return nil
}

// GetSubwayToken SubwayToken Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetSubwayToken() string {
	return r._subwayToken
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始日期,格式yyyy-mm-dd
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束日期,格式yyyy-mm-dd
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetSource is Source Setter
// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5, 汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetSource() string {
	return r._source
}

// SetSearchType is SearchType Setter
// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetSearchType(_searchType string) error {
	r._searchType = _searchType
	r.Set("search_type", _searchType)
	return nil
}

// GetSearchType SearchType Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetSearchType() string {
	return r._searchType
}

// SetCampaignId is CampaignId Setter
// 查询推广计划id
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaosimbarptcampadgroupbasegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarptcampadgroupbasegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
