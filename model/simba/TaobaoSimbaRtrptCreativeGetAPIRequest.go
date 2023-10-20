package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbartrptcreativegetAPIRequest 获取创意实时报表数据 API请求
// taobao.simba.rtrpt.creative.get
//
// 获取创意实时报表数据
type TaobaosimbartrptcreativegetAPIRequest struct {
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

// NewTaobaosimbartrptcreativegetRequest 初始化TaobaosimbartrptcreativegetAPIRequest对象
func NewTaobaosimbartrptcreativegetRequest() *TaobaosimbartrptcreativegetAPIRequest {
	return &TaobaosimbartrptcreativegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbartrptcreativegetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.creative.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbartrptcreativegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbartrptcreativegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户名
func (r *TaobaosimbartrptcreativegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbartrptcreativegetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaosimbartrptcreativegetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbartrptcreativegetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetCampaignId is CampaignId Setter
// 推广计划id
func (r *TaobaosimbartrptcreativegetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbartrptcreativegetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosimbartrptcreativegetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbartrptcreativegetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
