package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpProductStatusUpdateAPIRequest
修改P4P产品推广状态 API请求
alibaba.scbp.product.status.update

修改P4P产品推广状态 */
type AlibabaScbpProductStatusUpdateAPIRequest struct {
	model.Params
	// 产品ID列表
	_productIdList []int64
	// enabled:开启,disabled:暂停
	_status string
}

// New
