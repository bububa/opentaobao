package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCampaignOfflineLayeredfindAPIRequest 查询计划离线列表30天转化周期 API请求
// taobao.subway.campaign.offline.layeredfind
//
// 查询某计划离线列表
type TaobaoSubwayCampaignOfflineLayeredfindAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
	_pvTypeIn int64
	// 计划类型（直通车搜索-无线/pc：0；智能推广计划：8；销量明星计划：16；口碑L店计划：17；新享一键推广计划-独立结算账户(策略中心)：21；超级直播-一键推广计划（策略中心：订单模式、计划不复用：22；大快消一键推广计划（策略中心）：23；超级直播-持续推广计划（策略中心:计划模式、可复用、类似单品）：24；合约广告、流量卡计划：31；极速推计划：37；AI智投：38；）
	_campaignTypeNotIn int64
	// 需要查询的计划id，不传表示不限制
	_campaignIdEqual int64
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期30-30天
	_effect int64
}

// NewTaobaoSubwayCampaignOfflineLayeredfindRequest 初始化TaobaoSubwayCampaignOfflineLayeredfindAPIRequest对象
func NewTaobaoSubwayCampaignOfflineLayeredfindRequest() *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest {
	return &TaobaoSubwayCampaignOfflineLayeredfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.campaign.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetCampaignTypeNotIn is CampaignTypeNotIn Setter
// 计划类型（直通车搜索-无线/pc：0；智能推广计划：8；销量明星计划：16；口碑L店计划：17；新享一键推广计划-独立结算账户(策略中心)：21；超级直播-一键推广计划（策略中心：订单模式、计划不复用：22；大快消一键推广计划（策略中心）：23；超级直播-持续推广计划（策略中心:计划模式、可复用、类似单品）：24；合约广告、流量卡计划：31；极速推计划：37；AI智投：38；）
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetCampaignTypeNotIn(_campaignTypeNotIn int64) error {
	r._campaignTypeNotIn = _campaignTypeNotIn
	r.Set("campaign_type_not_in", _campaignTypeNotIn)
	return nil
}

// GetCampaignTypeNotIn CampaignTypeNotIn Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetCampaignTypeNotIn() int64 {
	return r._campaignTypeNotIn
}

// SetCampaignIdEqual is CampaignIdEqual Setter
// 需要查询的计划id，不传表示不限制
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetCampaignIdEqual(_campaignIdEqual int64) error {
	r._campaignIdEqual = _campaignIdEqual
	r.Set("campaign_id_equal", _campaignIdEqual)
	return nil
}

// GetCampaignIdEqual CampaignIdEqual Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetCampaignIdEqual() int64 {
	return r._campaignIdEqual
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayCampaignOfflineLayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}
