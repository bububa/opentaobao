package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreRefusedAPIRequest
Pos端门店拒单 API请求
taobao.omniorder.store.refused

ISV Pos端门店拒单，通知星盘 */
type TaobaoOmniorderStoreRefusedAPIRequest struct {
	model.Params
	// 淘宝交易主订单ID
	_tid int64
	// 子订单列表
	_subOrderList []SubOrder
	// ISV的系统时间
	_reportTimestamp int64
	// 跟踪Id
	_traceId string
}

// New
