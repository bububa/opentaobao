package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataSelfProductDateGetAPIRequest 获取客户产品相关表现数据的可用时间范围 API请求
// alibaba.mydata.self.product.date.get
//
// 获取客户产品相关表现数据的可用时间范围
type AlibabaMydataSelfProductDateGetAPIRequest struct {
	model.Params
	// 统计周期类型，可以为"day"，"week"，"month"
	_statisticsType string
}

// NewAlibabaMydataSelfProductDateGetRequest 初始化AlibabaMydataSelfProductDateGetAPIRequest对象
func NewAlibabaMydataSelfProductDateGetRequest() *AlibabaMydataSelfProductDateGetAPIRequest {
	return &AlibabaMydataSelfProductDateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMydataSelfProductDateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.self.product.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMydataSelfProductDateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMydataSelfProductDateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatisticsType is StatisticsType Setter
// 统计周期类型，可以为&#34;day&#34;，&#34;week&#34;，&#34;month&#34;
func (r *AlibabaMydataSelfProductDateGetAPIRequest) SetStatisticsType(_statisticsType string) error {
	r._statisticsType = _statisticsType
	r.Set("statistics_type", _statisticsType)
	return nil
}

// GetStatisticsType StatisticsType Getter
func (r AlibabaMydataSelfProductDateGetAPIRequest) GetStatisticsType() string {
	return r._statisticsType
}
