package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemBrandQueryAPIRequest
品牌信息查询 API请求
alibaba.wdk.item.brand.query

品牌信息查询 */
type AlibabaWdkItemBrandQueryAPIRequest struct {
	model.Params
	// 查询关键词，不填则查询全部
	_keyword string
	// 起始位置
	_offset int64
	// 一页大小
	_pageSize int64
}

// New
