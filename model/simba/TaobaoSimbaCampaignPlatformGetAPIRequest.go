package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignPlatformGetAPIRequest 取得一个推广计划的投放平台设置 API请求
// taobao.simba.campaign.platform.get
//
// 获得一个推广计划的投放平台设置
type TaobaoSimbaCampaignPlatformGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignPlatformGetRequest 初始化TaobaoSimbaCampaignPlatformGetAPIRequest对象
func NewTaobaoSimbaCampaignPlatformGetRequest() *TaobaoSimbaCampaignPlatformGetAPIRequest {
	return &TaobaoSimbaCampaignPlatformGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignPlatformGetAPIRequest) Reset() {
	r._nick = ""
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.platform.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignPlatformGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignPlatformGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignPlatformGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoSimbaCampaignPlatformGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignPlatformGetRequest()
	},
}

// GetTaobaoSimbaCampaignPlatformGetRequest 从 sync.Pool 获取 TaobaoSimbaCampaignPlatformGetAPIRequest
func GetTaobaoSimbaCampaignPlatformGetAPIRequest() *TaobaoSimbaCampaignPlatformGetAPIRequest {
	return poolTaobaoSimbaCampaignPlatformGetAPIRequest.Get().(*TaobaoSimbaCampaignPlatformGetAPIRequest)
}

// ReleaseTaobaoSimbaCampaignPlatformGetAPIRequest 将 TaobaoSimbaCampaignPlatformGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignPlatformGetAPIRequest(v *TaobaoSimbaCampaignPlatformGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignPlatformGetAPIRequest.Put(v)
}
