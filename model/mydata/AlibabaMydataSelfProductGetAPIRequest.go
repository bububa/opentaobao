package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamydataselfproductgetAPIRequest 获取客户产品相关表现数据 API请求
// alibaba.mydata.self.product.get
//
// 获取客户产品相关表现数据
type AlibabamydataselfproductgetAPIRequest struct {
	model.Params
	// 待查询产品id列表
	_productIds []int64
	// 统计周期，可以为"day", "week", "month"
	_statisticsType string
	// 统计日期
	_statDate string
}

// NewAlibabamydataselfproductgetRequest 初始化AlibabamydataselfproductgetAPIRequest对象
func NewAlibabamydataselfproductgetRequest() *AlibabamydataselfproductgetAPIRequest {
	return &AlibabamydataselfproductgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamydataselfproductgetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.self.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamydataselfproductgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamydataselfproductgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIds is ProductIds Setter
// 待查询产品id列表
func (r *AlibabamydataselfproductgetAPIRequest) SetProductIds(_productIds []int64) error {
	r._productIds = _productIds
	r.Set("product_ids", _productIds)
	return nil
}

// GetProductIds ProductIds Getter
func (r AlibabamydataselfproductgetAPIRequest) GetProductIds() []int64 {
	return r._productIds
}

// SetStatisticsType is StatisticsType Setter
// 统计周期，可以为&#34;day&#34;, &#34;week&#34;, &#34;month&#34;
func (r *AlibabamydataselfproductgetAPIRequest) SetStatisticsType(_statisticsType string) error {
	r._statisticsType = _statisticsType
	r.Set("statistics_type", _statisticsType)
	return nil
}

// GetStatisticsType StatisticsType Getter
func (r AlibabamydataselfproductgetAPIRequest) GetStatisticsType() string {
	return r._statisticsType
}

// SetStatDate is StatDate Setter
// 统计日期
func (r *AlibabamydataselfproductgetAPIRequest) SetStatDate(_statDate string) error {
	r._statDate = _statDate
	r.Set("stat_date", _statDate)
	return nil
}

// GetStatDate StatDate Getter
func (r AlibabamydataselfproductgetAPIRequest) GetStatDate() string {
	return r._statDate
}
