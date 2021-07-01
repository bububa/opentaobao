package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOrderQueryAPIRequest
大麦-查询分销单 API请求
alibaba.damai.maitix.order.query

查询分销单 */
type AlibabaDamaiMaitixOrderQueryAPIRequest struct {
	model.Params
	// 分销单查询入参
	_param *MoaOrderQueryParam
}

// New
