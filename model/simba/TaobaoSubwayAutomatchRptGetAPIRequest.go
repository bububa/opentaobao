package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayautomatchrptgetAPIRequest 查询流量智选天级报告 API请求
// taobao.subway.automatch.rpt.get
//
// 查询流量智选天级报告
type TaobaosubwayautomatchrptgetAPIRequest struct {
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

// NewTaobaosubwayautomatchrptgetRequest 初始化TaobaosubwayautomatchrptgetAPIRequest对象
func NewTaobaosubwayautomatchrptgetRequest() *TaobaosubwayautomatchrptgetAPIRequest {
	return &TaobaosubwayautomatchrptgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayautomatchrptgetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.automatch.rpt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayautomatchrptgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayautomatchrptgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosubwayautomatchrptgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwayautomatchrptgetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartDate is StartDate Setter
// 起始日期
func (r *TaobaosubwayautomatchrptgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaosubwayautomatchrptgetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 终止日期
func (r *TaobaosubwayautomatchrptgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaosubwayautomatchrptgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaosubwayautomatchrptgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosubwayautomatchrptgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosubwayautomatchrptgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosubwayautomatchrptgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
