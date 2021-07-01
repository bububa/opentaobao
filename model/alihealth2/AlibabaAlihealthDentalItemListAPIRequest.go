package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalItemListAPIRequest
ISV获取口腔标品列表 API请求
alibaba.alihealth.dental.item.list

ISV获取口腔标品列表 */
type AlibabaAlihealthDentalItemListAPIRequest struct {
	model.Params
	// 是否需要测试商品
	_needTestItem bool
}

// New
