package beehive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBeehiveItemCpsUrlAPIRequest
分佣链接生成接口 API请求
taobao.beehive.item.cps.url

传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接 */
type TaobaoBeehiveItemCpsUrlAPIRequest struct {
	model.Params
	// 平台，一般为手机
	_platform string
	// 达人ID
	_adUserId int64
	// 站外是1
	_sourceType int64
	// 业务方，新浪为sina
	_bizType string
	// 商品ID
	_itemId int64
}

// New
