package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayMarshLandRptGetAPIRequest 获取捡漏词包分时报表数据 API请求
// taobao.subway.marsh.land.rpt.get
//
// 获取捡漏词包分时报表数据
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
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) Reset() {
	r._nick = ""
	r._endDate = ""
	r._adgroupIdEqual = ""
	r._isAutoMatchEqual = ""
	r._campaignIdEqual = ""
	r._startDate = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.marsh.land.rpt.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetNick() string {
	return r._nick
}

// SetEndDate is EndDate Setter
// 2021-05-11
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetAdgroupIdEqual is AdgroupIdEqual Setter
// 推广组id
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetAdgroupIdEqual(_adgroupIdEqual string) error {
	r._adgroupIdEqual = _adgroupIdEqual
	r.Set("adgroup_id_equal", _adgroupIdEqual)
	return nil
}

// GetAdgroupIdEqual AdgroupIdEqual Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetAdgroupIdEqual() string {
	return r._adgroupIdEqual
}

// SetIsAutoMatchEqual is IsAutoMatchEqual Setter
// 词包类型（捡漏词包填19）
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetIsAutoMatchEqual(_isAutoMatchEqual string) error {
	r._isAutoMatchEqual = _isAutoMatchEqual
	r.Set("is_auto_match_equal", _isAutoMatchEqual)
	return nil
}

// GetIsAutoMatchEqual IsAutoMatchEqual Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetIsAutoMatchEqual() string {
	return r._isAutoMatchEqual
}

// SetCampaignIdEqual is CampaignIdEqual Setter
// 计划id
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetCampaignIdEqual(_campaignIdEqual string) error {
	r._campaignIdEqual = _campaignIdEqual
	r.Set("campaign_id_equal", _campaignIdEqual)
	return nil
}

// GetCampaignIdEqual CampaignIdEqual Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetCampaignIdEqual() string {
	return r._campaignIdEqual
}

// SetStartDate is StartDate Setter
// 2021-05-05
func (r *TaobaoSubwayMarshLandRptGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoSubwayMarshLandRptGetAPIRequest) GetStartDate() string {
	return r._startDate
}

var poolTaobaoSubwayMarshLandRptGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayMarshLandRptGetRequest()
	},
}

// GetTaobaoSubwayMarshLandRptGetRequest 从 sync.Pool 获取 TaobaoSubwayMarshLandRptGetAPIRequest
func GetTaobaoSubwayMarshLandRptGetAPIRequest() *TaobaoSubwayMarshLandRptGetAPIRequest {
	return poolTaobaoSubwayMarshLandRptGetAPIRequest.Get().(*TaobaoSubwayMarshLandRptGetAPIRequest)
}

// ReleaseTaobaoSubwayMarshLandRptGetAPIRequest 将 TaobaoSubwayMarshLandRptGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayMarshLandRptGetAPIRequest(v *TaobaoSubwayMarshLandRptGetAPIRequest) {
	v.Reset()
	poolTaobaoSubwayMarshLandRptGetAPIRequest.Put(v)
}
