package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobrandstartshoprptcampaigngetAPIRequest 明星店铺推广计划报表数据查询 API请求
// taobao.brand.startshop.rpt.campaign.get
//
// 获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaobrandstartshoprptcampaigngetAPIRequest struct {
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

// NewTaobaobrandstartshoprptcampaigngetRequest 初始化TaobaobrandstartshoprptcampaigngetAPIRequest对象
func NewTaobaobrandstartshoprptcampaigngetRequest() *TaobaobrandstartshoprptcampaigngetAPIRequest {
	return &TaobaobrandstartshoprptcampaigngetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobrandstartshoprptcampaigngetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.startshop.rpt.campaign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobrandstartshoprptcampaigngetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobrandstartshoprptcampaigngetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraffictype is Traffictype Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaobrandstartshoprptcampaigngetAPIRequest) SetTraffictype(_traffictype string) error {
	r._traffictype = _traffictype
	r.Set("traffictype", _traffictype)
	return nil
}

// GetTraffictype Traffictype Getter
func (r TaobaobrandstartshoprptcampaigngetAPIRequest) GetTraffictype() string {
	return r._traffictype
}

// SetStartdate is Startdate Setter
// 查询开始时间(最多查询90天数据)
func (r *TaobaobrandstartshoprptcampaigngetAPIRequest) SetStartdate(_startdate string) error {
	r._startdate = _startdate
	r.Set("startdate", _startdate)
	return nil
}

// GetStartdate Startdate Getter
func (r TaobaobrandstartshoprptcampaigngetAPIRequest) GetStartdate() string {
	return r._startdate
}

// SetEnddate is Enddate Setter
// 查询截至时间(最晚查询到昨天)
func (r *TaobaobrandstartshoprptcampaigngetAPIRequest) SetEnddate(_enddate string) error {
	r._enddate = _enddate
	r.Set("enddate", _enddate)
	return nil
}

// GetEnddate Enddate Getter
func (r TaobaobrandstartshoprptcampaigngetAPIRequest) GetEnddate() string {
	return r._enddate
}

// SetEffect is Effect Setter
// 转化周期,默认15天,可选 3,7,15
func (r *TaobaobrandstartshoprptcampaigngetAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaobrandstartshoprptcampaigngetAPIRequest) GetEffect() int64 {
	return r._effect
}
