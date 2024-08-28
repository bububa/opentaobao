package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTmsCutConfirm 配拦截失败CP确认结果并回告
// alibaba.alihealth.tms.cut.confirm
//
// 配拦截失败CP确认结果并回告
func AlibabaAlihealthTmsCutConfirm(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTmsCutConfirmAPIRequest, resp *alihealth2.AlibabaAlihealthTmsCutConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
