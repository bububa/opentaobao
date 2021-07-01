package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemsTagQueryAPIRequest
打标结果查询-商品维度 API请求
taobao.qimen.items.tag.query

调用该接口，查询某个/某批商品上的标 */
type TaobaoQimenItemsTagQueryAPIRequest struct {
	model.Params
	// 线上淘宝商品ID，long，必填
	_itemIds []int64
}

// New
