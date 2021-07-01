package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantskuQueryAPIRequest
商家商品信息查询 API请求
alibaba.wdk.item.merchantsku.query

商家商品信息查询 */
type AlibabaWdkItemMerchantskuQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 机构编码
	_orgCode string
}

// New
