package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantCategoryQueryAPIRequest
三江erp对接类目查询接口 API请求
alibaba.wdk.merchant.category.query

三江erp对接类目查询接口 */
type AlibabaWdkMerchantCategoryQueryAPIRequest struct {
	model.Params
	// 搜索关键词，可不填就查全部
	_keyword string
	// 类目起点，可不填从根目录开始查
	_rootCategoryCode string
}

// New
