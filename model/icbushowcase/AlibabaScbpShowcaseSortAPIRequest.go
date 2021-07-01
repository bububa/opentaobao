package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpShowcaseSortAPIRequest
橱窗顺序变更 API请求
alibaba.scbp.showcase.sort

橱窗顺序变更 */
type AlibabaScbpShowcaseSortAPIRequest struct {
	model.Params
	// 要移动的橱窗id
	_windowId int64
	// 当前位置（从1开始）
	_sourceOrder int64
	// 目标位置（从1开始）
	_targetOrder int64
}

// New
