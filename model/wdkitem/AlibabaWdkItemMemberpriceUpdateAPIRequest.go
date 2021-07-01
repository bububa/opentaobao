package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMemberpriceUpdateAPIRequest
商品售价会员价修改 API请求
alibaba.wdk.item.memberprice.update

商品售价会员价修改 */
type AlibabaWdkItemMemberpriceUpdateAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 商品编码
	_skuCode string
	// 售价，单位分，售价会员价至少填一个
	_skuPrice int64
	// 会员价，单位分
	_skuMemberPrice int64
	// 是否清空会员价
	_cleanSkuMemberPrice bool
	// 时间戳
	_timeStamp int64
}

// New
