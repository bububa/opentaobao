package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStartshopRptWordpackageGetAPIRequest 明星店铺品牌流量包报表数据查询 API请求
// taobao.brand.startshop.rpt.wordpackage.get
//
// 获取明星店铺广告词包分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStartshopRptWordpackageGetAPIRequest struct {
	model.Params
	// 开始日期
	_startDate string
	// 结束日期
	_endDate string
	// 转化周期
	_effect string
	// 每页显示条数(0,200]
	_pageSize string
	// 当前页数 ,从1开始
	_pageIndex string
	// 流量类型
	_trafficType string
}

// NewTaobaoBrandStartshopRptWordpackageGetRequest 初始化TaobaoBrandStartshopRptWordpackageGetAPIRequest对象
func NewTaobaoBrandStartshopRptWordpackageGetRequest() *TaobaoBrandStartshopRptWordpackageGetAPIRequest {
	return &TaobaoBrandStartshopRptWordpackageGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.startshop.rpt.wordpackage.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 开始日期
func (r *TaobaoBrandStartshopRptWordpackageGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *TaobaoBrandStartshopRptWordpackageGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetEffect is Effect Setter
// 转化周期
func (r *TaobaoBrandStartshopRptWordpackageGetAPIRequest) SetEffect(_effect string) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetEffect() string {
	return r._effect
}

// SetPageSize is PageSize Setter
// 每页显示条数(0,200]
func (r *TaobaoBrandStartshopRptWordpackageGetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 当前页数 ,从1开始
func (r *TaobaoBrandStartshopRptWordpackageGetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetTrafficType is TrafficType Setter
// 流量类型
func (r *TaobaoBrandStartshopRptWordpackageGetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaoBrandStartshopRptWordpackageGetAPIRequest) GetTrafficType() string {
	return r._trafficType
}
