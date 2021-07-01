package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTagItemsQueryAPIRequest
打标结果查询-标维度 API请求
taobao.qimen.tag.items.query

调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。 */
type TaobaoQimenTagItemsQueryAPIRequest struct {
	model.Params
	// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
	_tagType string
	// 备注，string（500）
	_remark string
}

// New
