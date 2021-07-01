package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenEntryorderConfirmAPIRequest
入库单确认接口 API请求
taobao.qimen.entryorder.confirm

WMS调用接口，回传入库单信息; */
type TaobaoQimenEntryorderConfirmAPIRequest struct {
	model.Params
	//
	_request *EntryOrderConfirmRequest
}

// New
