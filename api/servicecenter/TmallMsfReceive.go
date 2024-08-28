package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallMsfReceive 签收接口
// tmall.msf.receive
//
// 签收接口
func TmallMsfReceive(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallMsfReceiveAPIRequest, resp *servicecenter.TmallMsfReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
