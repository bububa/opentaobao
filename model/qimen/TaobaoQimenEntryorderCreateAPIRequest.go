package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenEntryorderCreateAPIRequest
入库单创建接口 API请求
taobao.qimen.entryorder.create

ERP调用接口，创建入库单; */
type TaobaoQimenEntryorderCreateAPIRequest struct {
	model.Params
	//
	_request *EntryOrderCreateRequest
}

// New
