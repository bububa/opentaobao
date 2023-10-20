package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktimegetAPIRequest 获得当前系统时间 API请求
// alibaba.wdk.time.get
//
// 获得当前系统时间
type AlibabawdktimegetAPIRequest struct {
	model.Params
}

// NewAlibabawdktimegetRequest 初始化AlibabawdktimegetAPIRequest对象
func NewAlibabawdktimegetRequest() *AlibabawdktimegetAPIRequest {
	return &AlibabawdktimegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktimegetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.time.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktimegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktimegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
