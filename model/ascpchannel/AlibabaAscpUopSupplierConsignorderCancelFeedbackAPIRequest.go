package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest
商家仓wms取消发货反馈回告服务 API请求
alibaba.ascp.uop.supplier.consignorder.cancel.feedback

履约单纬度通知商家仓wms取消发货结果反馈回告服务 */
type AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest struct {
	model.Params
	// 取消发货反馈回告请求
	_consignorderCancelFeedbackRequest *Consignordercancelfeedbackrequest
}

// New
