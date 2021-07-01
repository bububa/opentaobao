package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest
逆向销退入库单入库结果回告 API请求
alibaba.ascp.uop.supplier.reverseorder.instorage.feedback

ASCP按照逆向履约单纬度，通过该接口接收商家在退货完成时，自动创建销退单做入库回传。 */
type AlibabaAscpUopSupplierReverseorderInstorageFeedbackAPIRequest struct {
	model.Params
	// 销退单入库结果请求
	_instorageFeedbackRequest *Instoragefeedbackrequest
}

// New
