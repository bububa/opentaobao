package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTransferorderQueryAPIRequest
调拨单查询 API请求
taobao.qimen.transferorder.query

调拨单查询 */
type TaobaoQimenTransferorderQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenTransferorderQueryStruct
}

// New
