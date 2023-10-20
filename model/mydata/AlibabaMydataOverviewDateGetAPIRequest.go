package mydata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMydataOverviewDateGetAPIRequest 我的效果-获取数据周期 API请求
// alibaba.mydata.overview.date.get
//
// 获取数据管家我的效果API可以使用的数据周期
type AlibabaMydataOverviewDateGetAPIRequest struct {
	model.Params
}

// NewAlibabaMydataOverviewDateGetRequest 初始化AlibabaMydataOverviewDateGetAPIRequest对象
func NewAlibabaMydataOverviewDateGetRequest() *AlibabaMydataOverviewDateGetAPIRequest {
	return &AlibabaMydataOverviewDateGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMydataOverviewDateGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewDateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewDateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMydataOverviewDateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaMydataOverviewDateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMydataOverviewDateGetRequest()
	},
}

// GetAlibabaMydataOverviewDateGetRequest 从 sync.Pool 获取 AlibabaMydataOverviewDateGetAPIRequest
func GetAlibabaMydataOverviewDateGetAPIRequest() *AlibabaMydataOverviewDateGetAPIRequest {
	return poolAlibabaMydataOverviewDateGetAPIRequest.Get().(*AlibabaMydataOverviewDateGetAPIRequest)
}

// ReleaseAlibabaMydataOverviewDateGetAPIRequest 将 AlibabaMydataOverviewDateGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaMydataOverviewDateGetAPIRequest(v *AlibabaMydataOverviewDateGetAPIRequest) {
	v.Reset()
	poolAlibabaMydataOverviewDateGetAPIRequest.Put(v)
}
