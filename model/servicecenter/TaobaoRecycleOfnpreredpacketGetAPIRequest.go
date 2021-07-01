package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRecycleOfnpreredpacketGetAPIRequest
服务商查询前置补贴红包的最新数据 API请求
taobao.recycle.ofnpreredpacket.get

服务商查询前置补贴红包的最新数据 */
type TaobaoRecycleOfnpreredpacketGetAPIRequest struct {
	model.Params
	// 旧机单id
	_oldOrderId int64
}

// New
