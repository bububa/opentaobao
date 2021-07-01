package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingFullrangeQueryitemAPIRequest
全场活动查询换购品 API请求
alibaba.wdk.marketing.fullrange.queryitem

全场活动查询换购品 */
type AlibabaWdkMarketingFullrangeQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// New
