package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRxCaDoctorStatusSave ca认证获取医师认证结果
// alibaba.alihealth.rx.ca.doctor.status.save
//
// ca认证获取医师认证结果
func AlibabaAlihealthRxCaDoctorStatusSave(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRxCaDoctorStatusSaveAPIRequest, resp *alihealthoutflow.AlibabaAlihealthRxCaDoctorStatusSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
