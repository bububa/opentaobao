package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowPrescriptionUpdate 处方外流-修改处方
// alibaba.alihealth.outflow.prescription.update
//
// 阿里健康-处方外流-对外提供处方修改功能
func AlibabaAlihealthOutflowPrescriptionUpdate(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
