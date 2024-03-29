package simba

import (
	"net/url"
	"sync"

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
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 推广计划id
	_campaignId int64
	// 推广单元id
	_adgroupId int64
}

// NewTaobaoSimbaRtrptTargetingtagGetRequest 初始化TaobaoSimbaRtrptTargetingtagGetAPIRequest对象
func NewTaobaoSimbaRtrptTargetingtagGetRequest() *TaobaoSimbaRtrptTargetingtagGetAPIRequest {
	return &TaobaoSimbaRtrptTargetingtagGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) Reset() {
	r._nick = ""
	r._theDate = ""
	r._trafficType = ""
	r._campaignId = 0
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.targetingtag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 旺旺名称
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaRtrptTargetingtagGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaRtrptTargetingtagGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaRtrptTargetingtagGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaRtrptTargetingtagGetRequest()
	},
}

// GetTaobaoSimbaRtrptTargetingtagGetRequest 从 sync.Pool 获取 TaobaoSimbaRtrptTargetingtagGetAPIRequest
func GetTaobaoSimbaRtrptTargetingtagGetAPIRequest() *TaobaoSimbaRtrptTargetingtagGetAPIRequest {
	return poolTaobaoSimbaRtrptTargetingtagGetAPIRequest.Get().(*TaobaoSimbaRtrptTargetingtagGetAPIRequest)
}

// ReleaseTaobaoSimbaRtrptTargetingtagGetAPIRequest 将 TaobaoSimbaRtrptTargetingtagGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaRtrptTargetingtagGetAPIRequest(v *TaobaoSimbaRtrptTargetingtagGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaRtrptTargetingtagGetAPIRequest.Put(v)
}
