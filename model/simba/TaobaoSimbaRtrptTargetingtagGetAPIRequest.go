package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptTargetingtagGetAPIRequest 搜索人群实时报表 API请求
// taobao.simba.rtrpt.targetingtag.get
//
// 获取搜搜人群实时报表
type TaobaoSimbaRtrptTargetingtagGetAPIRequest struct {
	model.Params
	// 旺旺名称
	_nick string
	// 推广计划id
	_campaignId int64
	// 推广单元id
	_adgroupId int64
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
}

// NewTaobaoSimbaRtrptTargetingtagGetRequest 初始化TaobaoSimbaRtrptTargetingtagGetAPIRequest对象
func NewTaobaoSimbaRtrptTargetingtagGetRequest() *TaobaoSimbaRtrptTargetingtagGetAPIRequest {
	return &TaobaoSimbaRtrptTargetingtagGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.targetingtag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 旺旺名称
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// Set is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// Get TheDate Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// Set is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// Get TrafficType Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetTrafficType() string {
	return r._trafficType
}
