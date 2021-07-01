package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemsMarkingAPIRequest
商品通自动打标 API请求
taobao.qimen.items.marking

调用该接口，对商品进行XXXX标的打标、去标的动作。 */
type TaobaoQimenItemsMarkingAPIRequest struct {
	model.Params
	// 操作类型，string（50），ADD=打标，DELETE=去标，必填
	_actionType string
	// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
	_tagType string
	// 线上商品ID，long，必填
	_itemIds []int64
	// 备注，string（500）
	_remark string
}

// New
