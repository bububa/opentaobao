package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbartrptbidwordgetAPIRequest 获取推广词实时报表数据 API请求
// taobao.simba.rtrpt.bidword.get
//
// 获取推广词报表数据
type TaobaosimbartrptbidwordgetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
	// 推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
}

// NewTaobaosimbartrptbidwordgetRequest 初始化TaobaosimbartrptbidwordgetAPIRequest对象
func NewTaobaosimbartrptbidwordgetRequest() *TaobaosimbartrptbidwordgetAPIRequest {
	return &TaobaosimbartrptbidwordgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbartrptbidwordgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.bidword.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbartrptbidwordgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbartrptbidwordgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户名
func (r *TaobaosimbartrptbidwordgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbartrptbidwordgetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaosimbartrptbidwordgetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbartrptbidwordgetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaosimbartrptbidwordgetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbartrptbidwordgetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosimbartrptbidwordgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbartrptbidwordgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
