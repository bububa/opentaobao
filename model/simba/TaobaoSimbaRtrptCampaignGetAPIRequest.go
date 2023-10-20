package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbartrptcampaigngetAPIRequest 获取推广计划实时报表数据 API请求
// taobao.simba.rtrpt.campaign.get
//
// 获取推广计划实时报表数据
type TaobaosimbartrptcampaigngetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
}

// NewTaobaosimbartrptcampaigngetRequest 初始化TaobaosimbartrptcampaigngetAPIRequest对象
func NewTaobaosimbartrptcampaigngetRequest() *TaobaosimbartrptcampaigngetAPIRequest {
	return &TaobaosimbartrptcampaigngetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbartrptcampaigngetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbartrptcampaigngetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbartrptcampaigngetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户名
func (r *TaobaosimbartrptcampaigngetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbartrptcampaigngetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaosimbartrptcampaigngetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbartrptcampaigngetAPIRequest) GetTheDate() string {
	return r._theDate
}
