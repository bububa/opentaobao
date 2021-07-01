package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemPermitCheckAPIRequest
发品资质校验 API请求
taobao.item.permit.check

对淘宝商品发品、编辑前的预校验接口 */
type TaobaoItemPermitCheckAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 类目id
	_cid int64
	// 发布类型。可选值:fixed(一口价),auction(拍卖)
	_type string
}

// New
