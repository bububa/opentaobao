package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrRxQueryimageAPIRequest
o2o查看处方图片 API请求
alibaba.alihealth.nr.rx.queryimage

o2o商家查看处方图片，包括电子图片与纸质图片 */
type AlibabaAlihealthNrRxQueryimageAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
}

// New
