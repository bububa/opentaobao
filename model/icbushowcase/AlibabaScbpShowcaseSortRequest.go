package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
橱窗顺序变更 API请求
alibaba.scbp.showcase.sort

橱窗顺序变更
*/
type AlibabaScbpShowcaseSortRequest struct {
    model.Params
    // 要移动的橱窗id
    _windowId   int64
    // 当前位置（从1开始）
    _sourceOrder   int64
    // 目标位置（从1开始）
    _targetOrder   int64
}

// 初始化AlibabaScbpShowcaseSortRequest对象
func NewAlibabaScbpShowcaseSortRequest() *AlibabaScbpShowcaseSortRequest{
    return &AlibabaScbpShowcaseSortRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseSortRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.sort"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseSortRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WindowId Setter
// 要移动的橱窗id
func (r *AlibabaScbpShowcaseSortRequest) SetWindowId(_windowId int64) error {
    r._windowId = _windowId
    r.Set("window_id", _windowId)
    return nil
}

// WindowId Getter
func (r AlibabaScbpShowcaseSortRequest) GetWindowId() int64 {
    return r._windowId
}
// SourceOrder Setter
// 当前位置（从1开始）
func (r *AlibabaScbpShowcaseSortRequest) SetSourceOrder(_sourceOrder int64) error {
    r._sourceOrder = _sourceOrder
    r.Set("source_order", _sourceOrder)
    return nil
}

// SourceOrder Getter
func (r AlibabaScbpShowcaseSortRequest) GetSourceOrder() int64 {
    return r._sourceOrder
}
// TargetOrder Setter
// 目标位置（从1开始）
func (r *AlibabaScbpShowcaseSortRequest) SetTargetOrder(_targetOrder int64) error {
    r._targetOrder = _targetOrder
    r.Set("target_order", _targetOrder)
    return nil
}

// TargetOrder Getter
func (r AlibabaScbpShowcaseSortRequest) GetTargetOrder() int64 {
    return r._targetOrder
}
