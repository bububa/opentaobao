package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreAccpetedAPIRequest
Pos端门店接单接口 API请求
taobao.omniorder.store.accpeted

ISV Pos端门店接单，通知星盘 */
type TaobaoOmniorderStoreAccpetedAPIRequest struct {
	model.Params
	// 淘宝交易主订单ID
	_tid int64
	// 子订单列表
	_subOrderList []StoreAcceptedResult
	// ISV系统上报时间
	_reportTimestamp int64
	// 跟踪Id
	_traceId string
}

// New
