package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbartrpttargetingtaggetAPIRequest 搜索人群实时报表 API请求
// taobao.simba.rtrpt.targetingtag.get
//
// 获取搜搜人群实时报表
type TaobaosimbartrpttargetingtaggetAPIRequest struct {
	model.Params
	// 旺旺名称
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 推广计划id
	_campaignId int64
	// 推广单元id
	_adgroupId int64
}

// NewTaobaosimbartrpttargetingtaggetRequest 初始化TaobaosimbartrpttargetingtaggetAPIRequest对象
func NewTaobaosimbartrpttargetingtaggetRequest() *TaobaosimbartrpttargetingtaggetAPIRequest {
	return &TaobaosimbartrpttargetingtaggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.targetingtag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 旺旺名称
func (r *TaobaosimbartrpttargetingtaggetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaosimbartrpttargetingtaggetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaosimbartrpttargetingtaggetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaosimbartrpttargetingtaggetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbartrpttargetingtaggetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbartrpttargetingtaggetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
