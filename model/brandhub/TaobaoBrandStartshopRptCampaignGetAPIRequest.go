package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandStartshopRptCampaignGetAPIRequest
明星店铺推广计划报表数据查询 API请求
taobao.brand.startshop.rpt.campaign.get

获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandStartshopRptCampaignGetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_traffictype string
	// 查询开始时间(最多查询90天数据)
	_startdate string
	// 查询截至时间(最晚查询到昨天)
	_enddate string
	// 转化周期,默认15天,可选 3,7,15
	_effect int64
}

// NewTaobaoBrandStartshopRptCampaignGetRequest 初始化TaobaoBrandStartshopRptCampaignGetAPIRequest对象
func NewTaobaoBrandStartshopRptCampaignGetRequest() *TaobaoBrandStartshopRptCampaignGetAPIRequest {
	return &TaobaoBrandStartshopRptCampaignGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.startshop.rpt.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Traffictype Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetTraffictype(_traffictype string) error {
	r._traffictype = _traffictype
	r.Set("traffictype", _traffictype)
	return nil
}

// Get Traffictype Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetTraffictype() string {
	return r._traffictype
}

// Set is Startdate Setter
// 查询开始时间(最多查询90天数据)
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetStartdate(_startdate string) error {
	r._startdate = _startdate
	r.Set("startdate", _startdate)
	return nil
}

// Get Startdate Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetStartdate() string {
	return r._startdate
}

// Set is Enddate Setter
// 查询截至时间(最晚查询到昨天)
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetEnddate(_enddate string) error {
	r._enddate = _enddate
	r.Set("enddate", _enddate)
	return nil
}

// Get Enddate Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetEnddate() string {
	return r._enddate
}

// Set is Effect Setter
// 转化周期,默认15天,可选 3,7,15
func (r *TaobaoBrandStartshopRptCampaignGetAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// Get Effect Getter
func (r TaobaoBrandStartshopRptCampaignGetAPIRequest) GetEffect() int64 {
	return r._effect
}
