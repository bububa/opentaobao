package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeSpecialUsersMarkAPIRequest
专属下单可购买用户标记 API请求
taobao.opentrade.special.users.mark

对于专属下单的交易场景，根据openid标记用户可购买商品 */
type TaobaoOpentradeSpecialUsersMarkAPIRequest struct {
	model.Params
	// 是否目标用户，传入true后，用户可购买商品
	_hit bool
	// 本次待标记的用户列表，多个以逗号(,)分割，最大20个
	_openUserIds []string
	// 商品ID
	_itemId int64
	// 商品SKU ID，不存在传0
	_skuId int64
	// 用户状态，可任意传入，后续查询返回
	_status string
	// 单次购买最大限购数量
	_limitNum int64
}

// New
