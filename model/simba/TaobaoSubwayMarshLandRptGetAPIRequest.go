package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaymarshlandrptgetAPIRequest 获取捡漏词包分时报表数据 API请求
// taobao.subway.marsh.land.rpt.get
//
// 获取捡漏词包分时报表数据
type TaobaosubwaymarshlandrptgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 2021-05-11
	_endDate string
	// 推广组id
	_adgroupIdEqual string
	// 词包类型（捡漏词包填19）
	_isAutoMatchEqual string
	// 计划id
	_campaignIdEqual string
	// 2021-05-05
	_startDate string
}

// NewTaobaosubwaymarshlandrptgetRequest 初始化TaobaosubwaymarshlandrptgetAPIRequest对象
func NewTaobaosubwaymarshlandrptgetRequest() *TaobaosubwaymarshlandrptgetAPIRequest {
	return &TaobaosubwaymarshlandrptgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.marsh.land.rpt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosubwaymarshlandrptgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetNick() string {
	return r._nick
}

// SetEndDate is EndDate Setter
// 2021-05-11
func (r *TaobaosubwaymarshlandrptgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetAdgroupIdEqual is AdgroupIdEqual Setter
// 推广组id
func (r *TaobaosubwaymarshlandrptgetAPIRequest) SetAdgroupIdEqual(_adgroupIdEqual string) error {
	r._adgroupIdEqual = _adgroupIdEqual
	r.Set("adgroup_id_equal", _adgroupIdEqual)
	return nil
}

// GetAdgroupIdEqual AdgroupIdEqual Getter
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetAdgroupIdEqual() string {
	return r._adgroupIdEqual
}

// SetIsAutoMatchEqual is IsAutoMatchEqual Setter
// 词包类型（捡漏词包填19）
func (r *TaobaosubwaymarshlandrptgetAPIRequest) SetIsAutoMatchEqual(_isAutoMatchEqual string) error {
	r._isAutoMatchEqual = _isAutoMatchEqual
	r.Set("is_auto_match_equal", _isAutoMatchEqual)
	return nil
}

// GetIsAutoMatchEqual IsAutoMatchEqual Getter
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetIsAutoMatchEqual() string {
	return r._isAutoMatchEqual
}

// SetCampaignIdEqual is CampaignIdEqual Setter
// 计划id
func (r *TaobaosubwaymarshlandrptgetAPIRequest) SetCampaignIdEqual(_campaignIdEqual string) error {
	r._campaignIdEqual = _campaignIdEqual
	r.Set("campaign_id_equal", _campaignIdEqual)
	return nil
}

// GetCampaignIdEqual CampaignIdEqual Getter
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetCampaignIdEqual() string {
	return r._campaignIdEqual
}

// SetStartDate is StartDate Setter
// 2021-05-05
func (r *TaobaosubwaymarshlandrptgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaosubwaymarshlandrptgetAPIRequest) GetStartDate() string {
	return r._startDate
}
