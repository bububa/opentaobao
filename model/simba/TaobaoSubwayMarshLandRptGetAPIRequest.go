package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayMarshLandRptGetAPIRequest
获取捡漏词包分时报表数据 API请求
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据 */
type TaobaoSubwayMarshLandRptGetAPIRequest struct {
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

// NewTaobaoSubwayMarshLandRptGetRequest 初始化TaobaoSubwayMarshLandRptGetAPIRequest对象
func NewTaobaoSubwayMarshLandRptGetRequest() *TaobaoSubwayMarshLandRptGetAPIRequest {
	return &TaobaoSubwayMarshLandRptGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.marsh.land.rpt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is EndDate Setter
// 2021-05-11
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is AdgroupIdEqual Setter
// 推广组id
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetAdgroupIdEqual(_adgroupIdEqual string) error {
	r._adgroupIdEqual = _adgroupIdEqual
	r.Set("adgroup_id_equal", _adgroupIdEqual)
	return nil
}

// Get AdgroupIdEqual Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetAdgroupIdEqual() string {
	return r._adgroupIdEqual
}

// Set is IsAutoMatchEqual Setter
// 词包类型（捡漏词包填19）
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetIsAutoMatchEqual(_isAutoMatchEqual string) error {
	r._isAutoMatchEqual = _isAutoMatchEqual
	r.Set("is_auto_match_equal", _isAutoMatchEqual)
	return nil
}

// Get IsAutoMatchEqual Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetIsAutoMatchEqual() string {
	return r._isAutoMatchEqual
}

// Set is CampaignIdEqual Setter
// 计划id
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetCampaignIdEqual(_campaignIdEqual string) error {
	r._campaignIdEqual = _campaignIdEqual
	r.Set("campaign_id_equal", _campaignIdEqual)
	return nil
}

// Get CampaignIdEqual Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetCampaignIdEqual() string {
	return r._campaignIdEqual
}

// Set is StartDate Setter
// 2021-05-05
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetStartDate() string {
	return r._startDate
}
