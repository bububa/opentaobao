package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptTargetingtagGetAPIRequest 搜索人群离线报表 API请求
// taobao.simba.rpt.targetingtag.get
//
// 获取搜搜人群实时报表
type TaobaoSimbaRptTargetingtagGetAPIRequest struct {
	model.Params
	// 用户旺旺名称
	_nick string
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 推广计划id
	_campaignId int64
	// 推广单元id
	_adgroupId int64
}

// NewTaobaoSimbaRptTargetingtagGetRequest 初始化TaobaoSimbaRptTargetingtagGetAPIRequest对象
func NewTaobaoSimbaRptTargetingtagGetRequest() *TaobaoSimbaRptTargetingtagGetAPIRequest {
	return &TaobaoSimbaRptTargetingtagGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.targetingtag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 用户旺旺名称
func (r *TaobaoSimbaRptTargetingtagGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSimbaRptTargetingtagGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSimbaRptTargetingtagGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoSimbaRptTargetingtagGetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRptTargetingtagGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaRptTargetingtagGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaRptTargetingtagGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
