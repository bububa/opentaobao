package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallItemGetAPIRequest
获取商品详情物料 API请求
taobao.openmall.item.get

获取联盟开放的openmall商品 */
type TaobaoOpenmallItemGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// New
