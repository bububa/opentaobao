package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStartshopRptAccountGetAPIRequest 明星店铺账户报表数据查询 API请求
// taobao.brand.startshop.rpt.account.get
//
// 获取明星店铺广告主账户整体报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStartshopRptAccountGetAPIRequest struct {
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

// NewTaobaoBrandStartshopRptAccountGetRequest 初始化TaobaoBrandStartshopRptAccountGetAPIRequest对象
func NewTaobaoBrandStartshopRptAccountGetRequest() *TaobaoBrandStartshopRptAccountGetAPIRequest {
	return &TaobaoBrandStartshopRptAccountGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptAccountGetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.startshop.rpt.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptAccountGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBrandStartshopRptAccountGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStartshopRptAccountGetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaoBrandStartshopRptAccountGetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetEffect is Effect Setter
// 默认15天
func (r *TaobaoBrandStartshopRptAccountGetAPIRequest) SetEffect(_effect string) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoBrandStartshopRptAccountGetAPIRequest) GetEffect() string {
	return r._effect
}

// SetEndDate is EndDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaoBrandStartshopRptAccountGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoBrandStartshopRptAccountGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStartDate is StartDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaoBrandStartshopRptAccountGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoBrandStartshopRptAccountGetAPIRequest) GetStartDate() string {
	return r._startDate
}
