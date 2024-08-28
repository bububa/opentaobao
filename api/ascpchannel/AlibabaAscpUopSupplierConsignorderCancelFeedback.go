package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopSupplierConsignorderCancelFeedback 商家仓wms取消发货反馈回告服务
// alibaba.ascp.uop.supplier.consignorder.cancel.feedback
//
// 履约单纬度通知商家仓wms取消发货结果反馈回告服务
func AlibabaAscpUopSupplierConsignorderCancelFeedback(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIRequest, resp *ascpchannel.AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
