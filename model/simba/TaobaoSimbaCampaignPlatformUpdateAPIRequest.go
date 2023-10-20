package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignPlatformUpdateAPIRequest 更新一个推广计划的平台设置 API请求
// taobao.simba.campaign.platform.update
//
// 更新一个推广计划的平台设置
type TaobaoSimbaCampaignPlatformUpdateAPIRequest struct {
	model.Params
	// 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，8、无线站内搜索；16:无线站外搜索
	_searchChannels []string
	// 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。1、淘宝站内定向；2、站外定向；8、无线站内定向；16、无线站外定向
	_nonsearchChannels []string
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
	// 已经废弃
	_outsideDiscount int64
	// 已经废弃
	_mobileDiscount int64
}

// NewTaobaoSimbaCampaignPlatformUpdateRequest 初始化TaobaoSimbaCampaignPlatformUpdateAPIRequest对象
func NewTaobaoSimbaCampaignPlatformUpdateRequest() *TaobaoSimbaCampaignPlatformUpdateAPIRequest {
	return &TaobaoSimbaCampaignPlatformUpdateAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignPlatformUpdateAPIRequest) Reset() {
	r._searchChannels = r._searchChannels[:0]
	r._nonsearchChannels = r._nonsearchChannels[:0]
	r._nick = ""
	r._campaignId = 0
	r._outsideDiscount = 0
	r._mobileDiscount = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.platform.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchChannels is SearchChannels Setter
// 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，8、无线站内搜索；16:无线站外搜索
func (r *TaobaoSimbaCampaignPlatformUpdateAPIRequest) SetSearchChannels(_searchChannels []string) error {
	r._searchChannels = _searchChannels
	r.Set("search_channels", _searchChannels)
	return nil
}

// GetSearchChannels SearchChannels Getter
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetSearchChannels() []string {
	return r._searchChannels
}

// SetNonsearchChannels is NonsearchChannels Setter
// 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。1、淘宝站内定向；2、站外定向；8、无线站内定向；16、无线站外定向
func (r *TaobaoSimbaCampaignPlatformUpdateAPIRequest) SetNonsearchChannels(_nonsearchChannels []string) error {
	r._nonsearchChannels = _nonsearchChannels
	r.Set("nonsearch_channels", _nonsearchChannels)
	return nil
}

// GetNonsearchChannels NonsearchChannels Getter
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetNonsearchChannels() []string {
	return r._nonsearchChannels
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignPlatformUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignPlatformUpdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetOutsideDiscount is OutsideDiscount Setter
// 已经废弃
func (r *TaobaoSimbaCampaignPlatformUpdateAPIRequest) SetOutsideDiscount(_outsideDiscount int64) error {
	r._outsideDiscount = _outsideDiscount
	r.Set("outside_discount", _outsideDiscount)
	return nil
}

// GetOutsideDiscount OutsideDiscount Getter
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetOutsideDiscount() int64 {
	return r._outsideDiscount
}

// SetMobileDiscount is MobileDiscount Setter
// 已经废弃
func (r *TaobaoSimbaCampaignPlatformUpdateAPIRequest) SetMobileDiscount(_mobileDiscount int64) error {
	r._mobileDiscount = _mobileDiscount
	r.Set("mobile_discount", _mobileDiscount)
	return nil
}

// GetMobileDiscount MobileDiscount Getter
func (r TaobaoSimbaCampaignPlatformUpdateAPIRequest) GetMobileDiscount() int64 {
	return r._mobileDiscount
}

var poolTaobaoSimbaCampaignPlatformUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignPlatformUpdateRequest()
	},
}

// GetTaobaoSimbaCampaignPlatformUpdateRequest 从 sync.Pool 获取 TaobaoSimbaCampaignPlatformUpdateAPIRequest
func GetTaobaoSimbaCampaignPlatformUpdateAPIRequest() *TaobaoSimbaCampaignPlatformUpdateAPIRequest {
	return poolTaobaoSimbaCampaignPlatformUpdateAPIRequest.Get().(*TaobaoSimbaCampaignPlatformUpdateAPIRequest)
}

// ReleaseTaobaoSimbaCampaignPlatformUpdateAPIRequest 将 TaobaoSimbaCampaignPlatformUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignPlatformUpdateAPIRequest(v *TaobaoSimbaCampaignPlatformUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignPlatformUpdateAPIRequest.Put(v)
}
