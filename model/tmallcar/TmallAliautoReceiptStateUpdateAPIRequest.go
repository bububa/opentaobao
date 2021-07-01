package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAliautoReceiptStateUpdateAPIRequest
服务工单状态更新 API请求
tmall.aliauto.receipt.state.update

二轮车服务工单状态更新 */
type TmallAliautoReceiptStateUpdateAPIRequest struct {
	model.Params
	// FINISH:服务完成
	_status string
	// 服务工单id
	_receiptId int64
}

// New
