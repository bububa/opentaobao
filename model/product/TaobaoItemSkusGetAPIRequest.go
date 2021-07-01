package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemSkusGetAPIRequest
根据商品ID列表获取SKU信息 API请求
taobao.item.skus.get

* 获取多个商品下的所有sku
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
type TaobaoItemSkusGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
	_fields []string
	// sku所属商品数字id，必选。num_iid个数不能超过40个
	_numIids string
}

// New
