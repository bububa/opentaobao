package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobrandstartshoprptaccountgetAPIRequest 明星店铺账户报表数据查询 API请求
// taobao.brand.startshop.rpt.account.get
//
// 获取明星店铺广告主账户整体报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaobrandstartshoprptaccountgetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 默认15天
	_effect string
	// 开始时间(最多可查询最近90天)
	_endDate string
	// 截至时间(最晚到昨天)
	_startDate string
}

// NewTaobaobrandstartshoprptaccountgetRequest 初始化TaobaobrandstartshoprptaccountgetAPIRequest对象
func NewTaobaobrandstartshoprptaccountgetRequest() *TaobaobrandstartshoprptaccountgetAPIRequest {
	return &TaobaobrandstartshoprptaccountgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobrandstartshoprptaccountgetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.startshop.rpt.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobrandstartshoprptaccountgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobrandstartshoprptaccountgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaobrandstartshoprptaccountgetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaobrandstartshoprptaccountgetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetEffect is Effect Setter
// 默认15天
func (r *TaobaobrandstartshoprptaccountgetAPIRequest) SetEffect(_effect string) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaobrandstartshoprptaccountgetAPIRequest) GetEffect() string {
	return r._effect
}

// SetEndDate is EndDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaobrandstartshoprptaccountgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaobrandstartshoprptaccountgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStartDate is StartDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaobrandstartshoprptaccountgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaobrandstartshoprptaccountgetAPIRequest) GetStartDate() string {
	return r._startDate
}
