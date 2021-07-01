package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkShopQueryAPIRequest
门店查询接口 API请求
alibaba.wdk.shop.query

根据门店code查询门店信息 */
type AlibabaWdkShopQueryAPIRequest struct {
	model.Params
	// 如果不传，返回所有
	_ouCode string
}

// New
