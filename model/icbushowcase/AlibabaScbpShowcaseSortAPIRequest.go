package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpshowcasesortAPIRequest 橱窗顺序变更 API请求
// alibaba.scbp.showcase.sort
//
// 橱窗顺序变更
type AlibabascbpshowcasesortAPIRequest struct {
	model.Params
	// 要移动的橱窗id
	_windowId int64
	// 当前位置（从1开始）
	_sourceOrder int64
	// 目标位置（从1开始）
	_targetOrder int64
}

// NewAlibabascbpshowcasesortRequest 初始化AlibabascbpshowcasesortAPIRequest对象
func NewAlibabascbpshowcasesortRequest() *AlibabascbpshowcasesortAPIRequest {
	return &AlibabascbpshowcasesortAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpshowcasesortAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.sort"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpshowcasesortAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpshowcasesortAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWindowId is WindowId Setter
// 要移动的橱窗id
func (r *AlibabascbpshowcasesortAPIRequest) SetWindowId(_windowId int64) error {
	r._windowId = _windowId
	r.Set("window_id", _windowId)
	return nil
}

// GetWindowId WindowId Getter
func (r AlibabascbpshowcasesortAPIRequest) GetWindowId() int64 {
	return r._windowId
}

// SetSourceOrder is SourceOrder Setter
// 当前位置（从1开始）
func (r *AlibabascbpshowcasesortAPIRequest) SetSourceOrder(_sourceOrder int64) error {
	r._sourceOrder = _sourceOrder
	r.Set("source_order", _sourceOrder)
	return nil
}

// GetSourceOrder SourceOrder Getter
func (r AlibabascbpshowcasesortAPIRequest) GetSourceOrder() int64 {
	return r._sourceOrder
}

// SetTargetOrder is TargetOrder Setter
// 目标位置（从1开始）
func (r *AlibabascbpshowcasesortAPIRequest) SetTargetOrder(_targetOrder int64) error {
	r._targetOrder = _targetOrder
	r.Set("target_order", _targetOrder)
	return nil
}

// GetTargetOrder TargetOrder Getter
func (r AlibabascbpshowcasesortAPIRequest) GetTargetOrder() int64 {
	return r._targetOrder
}
