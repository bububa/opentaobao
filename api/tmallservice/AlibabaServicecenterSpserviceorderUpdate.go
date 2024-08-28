package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterSpserviceorderUpdate 服务供应链服务单更新
// alibaba.servicecenter.spserviceorder.update
//
// 服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
func AlibabaServicecenterSpserviceorderUpdate(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaServicecenterSpserviceorderUpdateAPIRequest, resp *tmallservice.AlibabaServicecenterSpserviceorderUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
