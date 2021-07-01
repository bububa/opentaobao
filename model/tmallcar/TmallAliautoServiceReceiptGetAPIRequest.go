package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAliautoServiceReceiptGetAPIRequest
isv查询服务工单详情 API请求
tmall.aliauto.service.receipt.get

isv查询服务工单详情 */
type TmallAliautoServiceReceiptGetAPIRequest struct {
	model.Params
	// 工单号
	_receiptId int64
}

// New
