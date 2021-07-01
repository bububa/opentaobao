package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTransferorderCreateAPIRequest
调拨单创建 API请求
taobao.qimen.transferorder.create

调拨单创建 */
type TaobaoQimenTransferorderCreateAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenTransferorderCreateStruct
}

// New
