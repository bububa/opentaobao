package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardConfirm 服务商确认服务完成
// tmall.servicecenter.workcard.confirm
//
// 提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
func TmallServicecenterWorkcardConfirm(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardConfirmAPIRequest, resp *tmallservice.TmallServicecenterWorkcardConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
