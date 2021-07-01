package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSkuGetAPIRequest
获取SKU API请求
taobao.item.sku.get

获取sku_id所对应的sku数据
sku_id对应的sku要属于传入的nick对应的卖家
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
type TaobaoItemSkuGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
	_fields string
	// Sku的id。可以通过taobao.item.seller.get得到
	_skuId int64
	// 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。
	_numIid int64
}

// New
