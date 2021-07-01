package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemPriceUpdateAPIRequest
五道口商品中心价格修改 API请求
alibaba.wdk.item.price.update

修改门店商品售价和会员价 */
type AlibabaWdkItemPriceUpdateAPIRequest struct {
	model.Params
	// 盒马门店id
	_storeId string
	// 商品编码
	_skuCode string
	// 价格，单位是分
	_skuPrice int64
	// 会员价格，单位是分，且不能大于价格
	_skuMemberPrice int64
}

// New
