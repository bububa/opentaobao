package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOrderCancelAPIRequest
大麦-库存释放 API请求
alibaba.damai.maitix.order.cancel

库存释放 */
type AlibabaDamaiMaitixOrderCancelAPIRequest struct {
	model.Params
	// 库存释放入参
	_param *MoaUnlockTicketParam
}

// New
