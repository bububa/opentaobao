package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleOrderPerformAPIRequest
回收订单履约V2 API请求
alibaba.idle.recycle.order.perform

闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化 */
type AlibabaIdleRecycleOrderPerformAPIRequest struct {
	model.Params
	// 参数
	_param0 *RecycleOrderSynDto
}

// New
