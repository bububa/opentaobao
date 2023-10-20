package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaGetAPIRequest 取得一个推广计划的投放地域设置 API请求
// taobao.simba.campaign.area.get
//
// 取得一个推广计划的投放地域设置
type TaobaoSimbaCampaignAreaGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignAreaGetRequest 初始化TaobaoSimbaCampaignAreaGetAPIRequest对象
func NewTaobaoSimbaCampaignAreaGetRequest() *TaobaoSimbaCampaignAreaGetAPIRequest {
	return &TaobaoSimbaCampaignAreaGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignAreaGetAPIRequest) Reset() {
	r._nick = ""
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.area.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAreaGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignAreaGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignAreaGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoSimbaCampaignAreaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignAreaGetRequest()
	},
}

// GetTaobaoSimbaCampaignAreaGetRequest 从 sync.Pool 获取 TaobaoSimbaCampaignAreaGetAPIRequest
func GetTaobaoSimbaCampaignAreaGetAPIRequest() *TaobaoSimbaCampaignAreaGetAPIRequest {
	return poolTaobaoSimbaCampaignAreaGetAPIRequest.Get().(*TaobaoSimbaCampaignAreaGetAPIRequest)
}

// ReleaseTaobaoSimbaCampaignAreaGetAPIRequest 将 TaobaoSimbaCampaignAreaGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignAreaGetAPIRequest(v *TaobaoSimbaCampaignAreaGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignAreaGetAPIRequest.Put(v)
}
