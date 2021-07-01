package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantBrandQueryAPIRequest
品牌查询接口 API请求
alibaba.wdk.merchant.brand.query

三江erp对接时，提供品牌查询的接口 */
type AlibabaWdkMerchantBrandQueryAPIRequest struct {
	model.Params
	// 关键词，不填就查全部
	_keyword string
	// 可不填，默认0
	_offset int64
	// 可不填，默认2000
	_pageSize int64
}

// New
