package mydata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataSelfProductGetAPIRequest 获取客户产品相关表现数据 API请求
// alibaba.mydata.self.product.get
//
// 获取客户产品相关表现数据
type AlibabaMydataSelfProductGetAPIRequest struct {
	model.Params
	// 待查询产品id列表
	_productIds []int64
	// 统计周期，可以为"day", "week", "month"
	_statisticsType string
	// 统计日期
	_statDate string
}

// NewAlibabaMydataSelfProductGetRequest 初始化AlibabaMydataSelfProductGetAPIRequest对象
func NewAlibabaMydataSelfProductGetRequest() *AlibabaMydataSelfProductGetAPIRequest {
	return &AlibabaMydataSelfProductGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMydataSelfProductGetAPIRequest) Reset() {
	r._productIds = r._productIds[:0]
	r._statisticsType = ""
	r._statDate = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMydataSelfProductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.self.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMydataSelfProductGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMydataSelfProductGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIds is ProductIds Setter
// 待查询产品id列表
func (r *AlibabaMydataSelfProductGetAPIRequest) SetProductIds(_productIds []int64) error {
	r._productIds = _productIds
	r.Set("product_ids", _productIds)
	return nil
}

// GetProductIds ProductIds Getter
func (r AlibabaMydataSelfProductGetAPIRequest) GetProductIds() []int64 {
	return r._productIds
}

// SetStatisticsType is StatisticsType Setter
// 统计周期，可以为&#34;day&#34;, &#34;week&#34;, &#34;month&#34;
func (r *AlibabaMydataSelfProductGetAPIRequest) SetStatisticsType(_statisticsType string) error {
	r._statisticsType = _statisticsType
	r.Set("statistics_type", _statisticsType)
	return nil
}

// GetStatisticsType StatisticsType Getter
func (r AlibabaMydataSelfProductGetAPIRequest) GetStatisticsType() string {
	return r._statisticsType
}

// SetStatDate is StatDate Setter
// 统计日期
func (r *AlibabaMydataSelfProductGetAPIRequest) SetStatDate(_statDate string) error {
	r._statDate = _statDate
	r.Set("stat_date", _statDate)
	return nil
}

// GetStatDate StatDate Getter
func (r AlibabaMydataSelfProductGetAPIRequest) GetStatDate() string {
	return r._statDate
}

var poolAlibabaMydataSelfProductGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMydataSelfProductGetRequest()
	},
}

// GetAlibabaMydataSelfProductGetRequest 从 sync.Pool 获取 AlibabaMydataSelfProductGetAPIRequest
func GetAlibabaMydataSelfProductGetAPIRequest() *AlibabaMydataSelfProductGetAPIRequest {
	return poolAlibabaMydataSelfProductGetAPIRequest.Get().(*AlibabaMydataSelfProductGetAPIRequest)
}

// ReleaseAlibabaMydataSelfProductGetAPIRequest 将 AlibabaMydataSelfProductGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaMydataSelfProductGetAPIRequest(v *AlibabaMydataSelfProductGetAPIRequest) {
	v.Reset()
	poolAlibabaMydataSelfProductGetAPIRequest.Put(v)
}
