package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMydataOverviewDateGetAPIRequest
我的效果-获取数据周期 API请求
alibaba.mydata.overview.date.get

获取数据管家我的效果API可以使用的数据周期 */
type AlibabaMydataOverviewDateGetAPIRequest struct {
	model.Params
}

// NewAlibabaMydataOverviewDateGetRequest 初始化AlibabaMydataOverviewDateGetAPIRequest对象
func NewAlibabaMydataOverviewDateGetRequest() *AlibabaMydataOverviewDateGetAPIRequest {
	return &AlibabaMydataOverviewDateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewDateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewDateGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
