package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOldposOrderCreateAPIRequest
淘鲜达外部商户老pos机产生的订单同步进淘鲜达 API请求
alibaba.wdk.oldpos.order.create

淘鲜达外部商户老pos机产生的订单同步进淘鲜达 */
type AlibabaWdkOldposOrderCreateAPIRequest struct {
	model.Params
	// 入参
	_posOrderCreateRequest *PosOrderCreateRequest
}

// New
