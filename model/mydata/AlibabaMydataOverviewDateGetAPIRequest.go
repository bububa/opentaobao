package mydata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamydataoverviewdategetAPIRequest 我的效果-获取数据周期 API请求
// alibaba.mydata.overview.date.get
//
// 获取数据管家我的效果API可以使用的数据周期
type AlibabamydataoverviewdategetAPIRequest struct {
	model.Params
}

// NewAlibabamydataoverviewdategetRequest 初始化AlibabamydataoverviewdategetAPIRequest对象
func NewAlibabamydataoverviewdategetRequest() *AlibabamydataoverviewdategetAPIRequest {
	return &AlibabamydataoverviewdategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamydataoverviewdategetAPIRequest) GetApiMethodName() string {
	return "alibaba.mydata.overview.date.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamydataoverviewdategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamydataoverviewdategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
