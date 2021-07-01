package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemCategoryQueryAPIRequest
类目查询接口 API请求
alibaba.wdk.item.category.query

类目查询接口 */
type AlibabaWdkItemCategoryQueryAPIRequest struct {
	model.Params
	// 查询关键词，不填查全部
	_keyword string
	// 从哪个类目开始查，不填从根类目开始查
	_rootCategoryCode string
}

// New
