package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest 浏览打点接口 API请求
// alibaba.alsc.growth.interactive.task.pageviewtrigger
//
// 浏览打点接口
type AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest struct {
	model.Params
	// 打点入参
	_pageViewRequest *PageViewRequest
}

// NewAlibabaalscgrowthinteractivetaskpageviewtriggerRequest 初始化AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest对象
func NewAlibabaalscgrowthinteractivetaskpageviewtriggerRequest() *AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest {
	return &AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.task.pageviewtrigger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageViewRequest is PageViewRequest Setter
// 打点入参
func (r *AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest) SetPageViewRequest(_pageViewRequest *PageViewRequest) error {
	r._pageViewRequest = _pageViewRequest
	r.Set("page_view_request", _pageViewRequest)
	return nil
}

// GetPageViewRequest PageViewRequest Getter
func (r AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest) GetPageViewRequest() *PageViewRequest {
	return r._pageViewRequest
}
