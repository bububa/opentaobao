package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSellerGetAPIRequest
获取单个商品详细信息 API请求
taobao.item.seller.get

获取单个商品的全部信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
type TaobaoItemSellerGetAPIRequest struct {
	model.Params
	// 需要返回的商品字段列表。可选值：Item商品结构体中所有字段均可返回，多个字段用“,”分隔。
	_fields string
	// 商品数字ID
	_numIid int64
}

// New
