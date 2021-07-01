package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenBatchpushticketAPIRequest
大麦换验平台-第三方对外开放-票单接口batchPushTicket API请求
alibaba.damai.mev.open.batchpushticket

批量推送票单 */
type AlibabaDamaiMevOpenBatchpushticketAPIRequest struct {
	model.Params
	// 入参thirdTicketSetOpenParamList
	_thirdTicketSetOpenParamList []ThirdTicketPushOpenParam
}

// New
