package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest
全场增加换购品 API请求
alibaba.wdk.marketing.fullrange.addexchangeitem

全场增加换购品 */
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemStairSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// New
