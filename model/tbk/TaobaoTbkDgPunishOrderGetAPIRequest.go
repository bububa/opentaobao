package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgPunishOrderGetAPIRequest
淘宝客-推广者-处罚订单查询 API请求
taobao.tbk.dg.punish.order.get

新增处罚订单查询API，提供媒体调用查询能力。这个是给媒体自己用的 */
type TaobaoTbkDgPunishOrderGetAPIRequest struct {
	model.Params
	// 入参的对象
	_afOrderOption *TopApiAfOrderOption
}

// New
