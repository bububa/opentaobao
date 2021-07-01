package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderAllocatedinfoSyncAPIRequest
分单结果同步给星盘 API请求
taobao.omniorder.allocatedinfo.sync

ISV分单完成，将分单结果同步给星盘 */
type TaobaoOmniorderAllocatedinfoSyncAPIRequest struct {
	model.Params
	// 淘宝交易主订单ID
	_tid int64
	// 分单状态，如： 等待中(Waiting)，已分单(Allocated)，分单失败(AllocateFail)
	_status string
	// 分单结果消息, 如果status为AllocateFail, 则表示失败的理由.
	_message string
	// 1231243213213
	_reportTimestamp int64
	// 门店的分单列表
	_subOrderList []StoreAllocatedResult
	// 跟踪Id
	_traceId string
}

// New
