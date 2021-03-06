package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseSortAPIRequest 橱窗顺序变更 API请求
// alibaba.scbp.showcase.sort
//
// 橱窗顺序变更
type AlibabaScbpShowcaseSortAPIRequest struct {
	model.Params
	// 要移动的橱窗id
	_windowId int64
	// 当前位置（从1开始）
	_sourceOrder int64
	// 目标位置（从1开始）
	_targetOrder int64
}

// NewAlibabaScbpShowcaseSortRequest 初始化AlibabaScbpShowcaseSortAPIRequest对象
func NewAlibabaScbpShowcaseSortRequest() *AlibabaScbpShowcaseSortAPIRequest {
	return &AlibabaScbpShowcaseSortAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseSortAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.sort"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseSortAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWindowId is WindowId Setter
// 要移动的橱窗id
func (r *AlibabaScbpShowcaseSortAPIRequest) SetWindowId(_windowId int64) error {
	r._windowId = _windowId
	r.Set("window_id", _windowId)
	return nil
}

// GetWindowId WindowId Getter
func (r AlibabaScbpShowcaseSortAPIRequest) GetWindowId() int64 {
	return r._windowId
}

// SetSourceOrder is SourceOrder Setter
// 当前位置（从1开始）
func (r *AlibabaScbpShowcaseSortAPIRequest) SetSourceOrder(_sourceOrder int64) error {
	r._sourceOrder = _sourceOrder
	r.Set("source_order", _sourceOrder)
	return nil
}

// GetSourceOrder SourceOrder Getter
func (r AlibabaScbpShowcaseSortAPIRequest) GetSourceOrder() int64 {
	return r._sourceOrder
}

// SetTargetOrder is TargetOrder Setter
// 目标位置（从1开始）
func (r *AlibabaScbpShowcaseSortAPIRequest) SetTargetOrder(_targetOrder int64) error {
	r._targetOrder = _targetOrder
	r.Set("target_order", _targetOrder)
	return nil
}

// GetTargetOrder TargetOrder Getter
func (r AlibabaScbpShowcaseSortAPIRequest) GetTargetOrder() int64 {
	return r._targetOrder
}
