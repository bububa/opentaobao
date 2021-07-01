package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuScrollQueryAPIRequest
门店商品批量游标方式查询接口 API请求
alibaba.wdk.sku.scroll.query

通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。 */
type AlibabaWdkSkuScrollQueryAPIRequest struct {
	model.Params
	// 商家类目编码
	_merchantCatCode string
	// 门店编码
	_ouCode string
	// 游标：第一次请求不用填写，否则请填写上一次请求返回的值，直到获取到足够的数据
	_scrollId string
	// 英文逗号分隔的商品编码，最多20个。如果配合门店字段使用，直接非游标方式返回商品数据
	_skuCodes string
}

// New
