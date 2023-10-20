package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbahourreportadgroupgetAPIRequest 推广单元小时级别实时报表查询 API请求
// taobao.simba.hour.report.adgroup.get
//
// 推广单元小时级别实时报表查询
type TaobaosimbahourreportadgroupgetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
	// 计划id
	_campaignId string
	// 推广单元id
	_adgroupId string
}

// NewTaobaosimbahourreportadgroupgetRequest 初始化TaobaosimbahourreportadgroupgetAPIRequest对象
func NewTaobaosimbahourreportadgroupgetRequest() *TaobaosimbahourreportadgroupgetAPIRequest {
	return &TaobaosimbahourreportadgroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.hour.report.adgroup.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbahourreportadgroupgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 时间
func (r *TaobaosimbahourreportadgroupgetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetTheDate() string {
	return r._theDate
}

// SetHour is Hour Setter
// 当前小时
func (r *TaobaosimbahourreportadgroupgetAPIRequest) SetHour(_hour string) error {
	r._hour = _hour
	r.Set("hour", _hour)
	return nil
}

// GetHour Hour Getter
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetHour() string {
	return r._hour
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *TaobaosimbahourreportadgroupgetAPIRequest) SetCampaignId(_campaignId string) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetCampaignId() string {
	return r._campaignId
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbahourreportadgroupgetAPIRequest) SetAdgroupId(_adgroupId string) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbahourreportadgroupgetAPIRequest) GetAdgroupId() string {
	return r._adgroupId
}
