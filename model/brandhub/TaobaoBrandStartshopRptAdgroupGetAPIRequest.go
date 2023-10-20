package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobrandstartshoprptadgroupgetAPIRequest 明星店铺推广单元报表数据查询 API请求
// taobao.brand.startshop.rpt.adgroup.get
//
// 获取明星店铺广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaobrandstartshoprptadgroupgetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 开始时间(最多可查询最近90天)
	_startDate string
	// 截至时间(最晚到昨天)
	_endDate string
	// 当前页数
	_pageIndex string
	// 每页条数
	_pageSize string
	// 转化周期默认15天,3,7,15
	_effect int64
}

// NewTaobaobrandstartshoprptadgroupgetRequest 初始化TaobaobrandstartshoprptadgroupgetAPIRequest对象
func NewTaobaobrandstartshoprptadgroupgetRequest() *TaobaobrandstartshoprptadgroupgetAPIRequest {
	return &TaobaobrandstartshoprptadgroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.startshop.rpt.adgroup.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaobrandstartshoprptadgroupgetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetStartDate is StartDate Setter
// 开始时间(最多可查询最近90天)
func (r *TaobaobrandstartshoprptadgroupgetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 截至时间(最晚到昨天)
func (r *TaobaobrandstartshoprptadgroupgetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TaobaobrandstartshoprptadgroupgetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TaobaobrandstartshoprptadgroupgetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期默认15天,3,7,15
func (r *TaobaobrandstartshoprptadgroupgetAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaobrandstartshoprptadgroupgetAPIRequest) GetEffect() int64 {
	return r._effect
}
