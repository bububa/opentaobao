package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemStoreskuUpdateAPIRequest
五道口商品中心门店商品修改 API请求
alibaba.wdk.item.storesku.update

五道口商品中心门店商品修改 */
type AlibabaWdkItemStoreskuUpdateAPIRequest struct {
	model.Params
	// 盒马门店id
	_storeId string
	// 商品编码
	_skuCode string
	// 1-可售   0-不可售
	_saleFlag int64
}

// New
