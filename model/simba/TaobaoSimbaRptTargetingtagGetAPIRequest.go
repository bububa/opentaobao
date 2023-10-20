package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarpttargetingtaggetAPIRequest 搜索人群离线报表 API请求
// taobao.simba.rpt.targetingtag.get
//
// 获取搜搜人群实时报表
type TaobaosimbarpttargetingtaggetAPIRequest struct {
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

// NewTaobaosimbarpttargetingtaggetRequest 初始化TaobaosimbarpttargetingtaggetAPIRequest对象
func NewTaobaosimbarpttargetingtaggetRequest() *TaobaosimbarpttargetingtaggetAPIRequest {
	return &TaobaosimbarpttargetingtaggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rpt.targetingtag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户旺旺名称
func (r *TaobaosimbarpttargetingtaggetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaosimbarpttargetingtaggetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaosimbarpttargetingtaggetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaosimbarpttargetingtaggetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaosimbarpttargetingtaggetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbarpttargetingtaggetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbarpttargetingtaggetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
