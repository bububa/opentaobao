package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarpttargetingtagbasegetAPIRequest 定向基础报表 API请求
// taobao.simba.rpt.targetingtagbase.get
//
// 获取定向基础报表
type TaobaosimbarpttargetingtagbasegetAPIRequest struct {
	model.Params
	// 被操作者昵称
	_nick string
	// 起始时间
	_startTime string
	// 结束时间
	_endTime string
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 分页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaosimbarpttargetingtagbasegetRequest 初始化TaobaosimbarpttargetingtagbasegetAPIRequest对象
func NewTaobaosimbarpttargetingtagbasegetRequest() *TaobaosimbarpttargetingtagbasegetAPIRequest {
	return &TaobaosimbarpttargetingtagbasegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.targetingtagbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者昵称
func (r *TaobaosimbarpttargetingtagbasegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 起始时间
func (r *TaobaosimbarpttargetingtagbasegetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaosimbarpttargetingtagbasegetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaosimbarpttargetingtagbasegetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosimbarpttargetingtagbasegetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaosimbarpttargetingtagbasegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TaobaosimbarpttargetingtagbasegetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaosimbarpttargetingtagbasegetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
