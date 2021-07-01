package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpShowcaseListAPIRequest
橱窗查询 API请求
alibaba.scbp.showcase.list

橱窗查询 */
type AlibabaScbpShowcaseListAPIRequest struct {
	model.Params
	// 每页展示个数
	_perPageSize int64
	// 页码
	_toPage int64
}

// New
