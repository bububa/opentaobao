package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletePaperformatAPIRequest
大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat API请求
alibaba.damai.mev.open.delete.paperformat

deletePaperFormat */
type AlibabaDamaiMevOpenDeletePaperformatAPIRequest struct {
	model.Params
	// 入参deletePaperFormatParam
	_deletePaperFormatParam *TicketPaperFormatIdOpenParam
}

// New
