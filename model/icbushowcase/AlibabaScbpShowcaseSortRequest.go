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
    windowId   int64
    // 当前位置（从1开始）
    sourceOrder   int64
    // 目标位置（从1开始）
    targetOrder   int64
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
func (r *AlibabaScbpShowcaseSortRequest) SetWindowId(windowId int64) error {
    r.windowId = windowId
    r.Set("window_id", windowId)
    return nil
}

// WindowId Getter
func (r AlibabaScbpShowcaseSortRequest) GetWindowId() int64 {
    return r.windowId
}
// SourceOrder Setter
// 当前位置（从1开始）
func (r *AlibabaScbpShowcaseSortRequest) SetSourceOrder(sourceOrder int64) error {
    r.sourceOrder = sourceOrder
    r.Set("source_order", sourceOrder)
    return nil
}

// SourceOrder Getter
func (r AlibabaScbpShowcaseSortRequest) GetSourceOrder() int64 {
    return r.sourceOrder
}
// TargetOrder Setter
// 目标位置（从1开始）
func (r *AlibabaScbpShowcaseSortRequest) SetTargetOrder(targetOrder int64) error {
    r.targetOrder = targetOrder
    r.Set("target_order", targetOrder)
    return nil
}

// TargetOrder Getter
func (r AlibabaScbpShowcaseSortRequest) GetTargetOrder() int64 {
    return r.targetOrder
}
