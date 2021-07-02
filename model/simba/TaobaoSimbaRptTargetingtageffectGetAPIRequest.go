package simba

import (
	"net/url"

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
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 起始时间
	_startTime string
	// 终止时间 ,必须小于今天
	_endTime string
	// 页面大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaoSimbaRptTargetingtageffectGetRequest 初始化TaobaoSimbaRptTargetingtageffectGetAPIRequest对象
func NewTaobaoSimbaRptTargetingtageffectGetRequest() *TaobaoSimbaRptTargetingtageffectGetAPIRequest {
	return &TaobaoSimbaRptTargetingtageffectGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.targetingtageffect.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 被操作者昵称
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 计划id
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// Set is StartTime Setter
// 起始时间
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 终止时间 ,必须小于今天
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is PageSize Setter
// 页面大小
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNumber Setter
// 页码
func (r *TaobaoSimbaRptTargetingtageffectGetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// Get PageNumber Getter
func (r TaobaoSimbaRptTargetingtageffectGetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
