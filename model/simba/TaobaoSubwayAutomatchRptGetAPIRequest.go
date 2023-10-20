package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayAutomatchRptGetAPIRequest 查询流量智选天级报告 API请求
// taobao.subway.automatch.rpt.get
//
// 查询流量智选天级报告
type TaobaoSubwayAutomatchRptGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 起始日期
	_startDate string
	// 终止日期
	_endDate string
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
}

// NewTaobaoSubwayAutomatchRptGetRequest 初始化TaobaoSubwayAutomatchRptGetAPIRequest对象
func NewTaobaoSubwayAutomatchRptGetRequest() *TaobaoSubwayAutomatchRptGetAPIRequest {
	return &TaobaoSubwayAutomatchRptGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) Reset() {
	r._nick = ""
	r._startDate = ""
	r._endDate = ""
	r._campaignId = 0
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.automatch.rpt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartDate is StartDate Setter
// 起始日期
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 终止日期
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSubwayAutomatchRptGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayAutomatchRptGetRequest()
	},
}

// GetTaobaoSubwayAutomatchRptGetRequest 从 sync.Pool 获取 TaobaoSubwayAutomatchRptGetAPIRequest
func GetTaobaoSubwayAutomatchRptGetAPIRequest() *TaobaoSubwayAutomatchRptGetAPIRequest {
	return poolTaobaoSubwayAutomatchRptGetAPIRequest.Get().(*TaobaoSubwayAutomatchRptGetAPIRequest)
}

// ReleaseTaobaoSubwayAutomatchRptGetAPIRequest 将 TaobaoSubwayAutomatchRptGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayAutomatchRptGetAPIRequest(v *TaobaoSubwayAutomatchRptGetAPIRequest) {
	v.Reset()
	poolTaobaoSubwayAutomatchRptGetAPIRequest.Put(v)
}
