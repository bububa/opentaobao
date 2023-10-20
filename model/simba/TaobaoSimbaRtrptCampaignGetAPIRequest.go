package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptCampaignGetAPIRequest 获取推广计划实时报表数据 API请求
// taobao.simba.rtrpt.campaign.get
//
// 获取推广计划实时报表数据
type TaobaoSimbaRtrptCampaignGetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
}

// NewTaobaoSimbaRtrptCampaignGetRequest 初始化TaobaoSimbaRtrptCampaignGetAPIRequest对象
func NewTaobaoSimbaRtrptCampaignGetRequest() *TaobaoSimbaRtrptCampaignGetAPIRequest {
	return &TaobaoSimbaRtrptCampaignGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptCampaignGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCampaignGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetTheDate() string {
	return r._theDate
}
