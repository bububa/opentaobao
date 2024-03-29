package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptTargetingtagbaseGetAPIRequest 定向基础报表 API请求
// taobao.simba.rpt.targetingtagbase.get
//
// 获取定向基础报表
type TaobaoSimbaRptTargetingtagbaseGetAPIRequest struct {
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

// NewTaobaoSimbaRptTargetingtagbaseGetRequest 初始化TaobaoSimbaRptTargetingtagbaseGetAPIRequest对象
func NewTaobaoSimbaRptTargetingtagbaseGetRequest() *TaobaoSimbaRptTargetingtagbaseGetAPIRequest {
	return &TaobaoSimbaRptTargetingtagbaseGetAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) Reset() {
	r._nick = ""
	r._startTime = ""
	r._endTime = ""
	r._campaignId = 0
	r._adgroupId = 0
	r._pageSize = 0
	r._pageNumber = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.targetingtagbase.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者昵称
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaoSimbaRptTargetingtagbaseGetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}

var poolTaobaoSimbaRptTargetingtagbaseGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaRptTargetingtagbaseGetRequest()
	},
}

// GetTaobaoSimbaRptTargetingtagbaseGetRequest 从 sync.Pool 获取 TaobaoSimbaRptTargetingtagbaseGetAPIRequest
func GetTaobaoSimbaRptTargetingtagbaseGetAPIRequest() *TaobaoSimbaRptTargetingtagbaseGetAPIRequest {
	return poolTaobaoSimbaRptTargetingtagbaseGetAPIRequest.Get().(*TaobaoSimbaRptTargetingtagbaseGetAPIRequest)
}

// ReleaseTaobaoSimbaRptTargetingtagbaseGetAPIRequest 将 TaobaoSimbaRptTargetingtagbaseGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaRptTargetingtagbaseGetAPIRequest(v *TaobaoSimbaRptTargetingtagbaseGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaRptTargetingtagbaseGetAPIRequest.Put(v)
}
