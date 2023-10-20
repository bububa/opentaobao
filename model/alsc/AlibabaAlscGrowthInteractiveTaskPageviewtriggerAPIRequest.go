package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest 浏览打点接口 API请求
// alibaba.alsc.growth.interactive.task.pageviewtrigger
//
// 浏览打点接口
type AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest struct {
	model.Params
	// 打点入参
	_pageViewRequest *PageViewRequest
}

// NewAlibabaAlscGrowthInteractiveTaskPageviewtriggerRequest 初始化AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest对象
func NewAlibabaAlscGrowthInteractiveTaskPageviewtriggerRequest() *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest {
	return &AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest) Reset() {
	r._pageViewRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.task.pageviewtrigger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageViewRequest is PageViewRequest Setter
// 打点入参
func (r *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest) SetPageViewRequest(_pageViewRequest *PageViewRequest) error {
	r._pageViewRequest = _pageViewRequest
	r.Set("page_view_request", _pageViewRequest)
	return nil
}

// GetPageViewRequest PageViewRequest Getter
func (r AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest) GetPageViewRequest() *PageViewRequest {
	return r._pageViewRequest
}

var poolAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscGrowthInteractiveTaskPageviewtriggerRequest()
	},
}

// GetAlibabaAlscGrowthInteractiveTaskPageviewtriggerRequest 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest
func GetAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest() *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest {
	return poolAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest.Get().(*AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest)
}

// ReleaseAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest 将 AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest(v *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIRequest.Put(v)
}
