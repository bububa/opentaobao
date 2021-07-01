package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayAutomatchRptGetAPIRequest
查询流量智选天级报告 API请求
taobao.subway.automatch.rpt.get

查询流量智选天级报告 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.automatch.rpt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is StartDate Setter
// 起始日期
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 终止日期
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is CampaignId Setter
// 计划id
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is AdgroupId Setter
// 推广组id
func (r *TaobaoSubwayAutomatchRptGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSubwayAutomatchRptGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
