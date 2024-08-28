package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderEpocReceive 电子保单数据接口
// tmall.servicecenter.spserviceorder.epoc.receive
//
// 电子保单数据回传接口
func TmallServicecenterSpserviceorderEpocReceive(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderEpocReceiveAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderEpocReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
