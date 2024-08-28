package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TmallTmjlappSapServiceorderCancel 取消售后服务单
// tmall.tmjlapp.sap.serviceorder.cancel
//
// SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
func TmallTmjlappSapServiceorderCancel(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TmallTmjlappSapServiceorderCancelAPIRequest, resp *tmallgenie.TmallTmjlappSapServiceorderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
