package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenEntryorderQueryAPIRequest
入库单查询接口 API请求
taobao.qimen.entryorder.query

ERP调用接口，查询入库单信息; */
type TaobaoQimenEntryorderQueryAPIRequest struct {
	model.Params
	//
	_request *EntryOrderQueryRequest
}

// New
