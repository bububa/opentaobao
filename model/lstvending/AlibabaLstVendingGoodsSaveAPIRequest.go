package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingGoodsSaveAPIRequest
自动售卖机商品回传 API请求
alibaba.lst.vending.goods.save

零售通自动售卖机商品数据回流。 */
type AlibabaLstVendingGoodsSaveAPIRequest struct {
	model.Params
	// 商品信息
	_goodsDTOList []VendingGoodsDto
}

// New
