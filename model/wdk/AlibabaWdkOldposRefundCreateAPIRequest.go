package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOldposRefundCreateAPIRequest
五道口外部商户老pos机产生的退款单同步进盒马 API请求
alibaba.wdk.oldpos.refund.create

淘鲜达外部商户老pos机产生的退款单同步进淘鲜达 */
type AlibabaWdkOldposRefundCreateAPIRequest struct {
	model.Params
	// 入参
	_posRefundCreateRequest *PosRefundCreateRequest
}

// New
