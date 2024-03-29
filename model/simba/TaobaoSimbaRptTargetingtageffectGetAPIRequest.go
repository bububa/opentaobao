package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptTargetingtageffectGetAPIRequest 获取定向效果报表数据 API请求
// taobao.simba.rpt.targetingtageffect.get
//
// 获取定向效果报表数据
type TaobaoSimbaRptTargetingtageffectGetAPIRequest struct {
	model.Params
	// 被操作者昵称
	_nick string
	// 起始时间
	_startTime string
	// 终止时间 ,必须小于今天
	_endTime string
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 页面大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaoSimbaRptTargetingtageffectGetRequest 初始化TaobaoSimbaRptTargetingtageffectGetAPIRequest对象
func NewTaobaoSimbaRptTargetingtageffectGetRequest() *TaobaoSimbaRptTargetingtageffectGetAPIRequest {
	return &TaobaoSimbaRptTargetingtageffectGetAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) Reset() {
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
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.targetingtageffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者昵称
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 终止时间 ,必须小于今天
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}

var poolTaobaoSimbaRptTargetingtageffectGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaRptTargetingtageffectGetRequest()
	},
}

// GetTaobaoSimbaRptTargetingtageffectGetRequest 从 sync.Pool 获取 TaobaoSimbaRptTargetingtageffectGetAPIRequest
func GetTaobaoSimbaRptTargetingtageffectGetAPIRequest() *TaobaoSimbaRptTargetingtageffectGetAPIRequest {
	return poolTaobaoSimbaRptTargetingtageffectGetAPIRequest.Get().(*TaobaoSimbaRptTargetingtageffectGetAPIRequest)
}

// ReleaseTaobaoSimbaRptTargetingtageffectGetAPIRequest 将 TaobaoSimbaRptTargetingtageffectGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaRptTargetingtageffectGetAPIRequest(v *TaobaoSimbaRptTargetingtageffectGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaRptTargetingtageffectGetAPIRequest.Put(v)
}
