package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOrderDirectrefundAPIRequest
大麦-直接退票 API请求
alibaba.damai.maitix.order.directrefund

大麦-退票 */
type AlibabaDamaiMaitixOrderDirectrefundAPIRequest struct {
	model.Params
	// 退票入参
	_param *MoaRefundAuditParam
}

// New
