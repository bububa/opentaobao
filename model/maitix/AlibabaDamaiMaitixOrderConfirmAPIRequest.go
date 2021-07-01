package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOrderConfirmAPIRequest
大麦-出票 API请求
alibaba.damai.maitix.order.confirm

出票 */
type AlibabaDamaiMaitixOrderConfirmAPIRequest struct {
	model.Params
	// 出票入参
	_param *MoaConfirmOrderParam
}

// New
