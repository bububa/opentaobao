package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowVisitinfoSync 处方外流-问诊、处方同步
// alibaba.alihealth.outflow.visitinfo.sync
//
// 阿里健康-处方外流-对外提供问诊、处方功能
func AlibabaAlihealthOutflowVisitinfoSync(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowVisitinfoSyncAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowVisitinfoSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
