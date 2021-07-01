package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenReturnorderConfirmAPIRequest
退货入库单确认接口 API请求
taobao.qimen.returnorder.confirm

taobao.qimen.returnorder.confirm */
type TaobaoQimenReturnorderConfirmAPIRequest struct {
	model.Params
	//
	_request *ReturnOrderConfirmRequest
}

// New
