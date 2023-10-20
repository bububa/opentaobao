package icbushowcase

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpShowcaseSortAPIRequest) Reset() {
	r._windowId = 0
	r._sourceOrder = 0
	r._targetOrder = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseSortAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.sort"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseSortAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpShowcaseSortAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaScbpShowcaseSortAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpShowcaseSortRequest()
	},
}

// GetAlibabaScbpShowcaseSortRequest 从 sync.Pool 获取 AlibabaScbpShowcaseSortAPIRequest
func GetAlibabaScbpShowcaseSortAPIRequest() *AlibabaScbpShowcaseSortAPIRequest {
	return poolAlibabaScbpShowcaseSortAPIRequest.Get().(*AlibabaScbpShowcaseSortAPIRequest)
}

// ReleaseAlibabaScbpShowcaseSortAPIRequest 将 AlibabaScbpShowcaseSortAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpShowcaseSortAPIRequest(v *AlibabaScbpShowcaseSortAPIRequest) {
	v.Reset()
	poolAlibabaScbpShowcaseSortAPIRequest.Put(v)
}
