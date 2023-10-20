package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaUpdateAPIRequest 更新一个推广计划的投放地域 API请求
// taobao.simba.campaign.area.update
//
// 更新一个推广计划的投放地域
type TaobaoSimbaCampaignAreaUpdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
	_area string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignAreaUpdateRequest 初始化TaobaoSimbaCampaignAreaUpdateAPIRequest对象
func NewTaobaoSimbaCampaignAreaUpdateRequest() *TaobaoSimbaCampaignAreaUpdateAPIRequest {
	return &TaobaoSimbaCampaignAreaUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignAreaUpdateAPIRequest) Reset() {
	r._nick = ""
	r._area = ""
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAreaUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.area.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAreaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignAreaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAreaUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignAreaUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetArea is Area Setter
// 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
func (r *TaobaoSimbaCampaignAreaUpdateAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r TaobaoSimbaCampaignAreaUpdateAPIRequest) GetArea() string {
	return r._area
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignAreaUpdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignAreaUpdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoSimbaCampaignAreaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignAreaUpdateRequest()
	},
}

// GetTaobaoSimbaCampaignAreaUpdateRequest 从 sync.Pool 获取 TaobaoSimbaCampaignAreaUpdateAPIRequest
func GetTaobaoSimbaCampaignAreaUpdateAPIRequest() *TaobaoSimbaCampaignAreaUpdateAPIRequest {
	return poolTaobaoSimbaCampaignAreaUpdateAPIRequest.Get().(*TaobaoSimbaCampaignAreaUpdateAPIRequest)
}

// ReleaseTaobaoSimbaCampaignAreaUpdateAPIRequest 将 TaobaoSimbaCampaignAreaUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignAreaUpdateAPIRequest(v *TaobaoSimbaCampaignAreaUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignAreaUpdateAPIRequest.Put(v)
}
