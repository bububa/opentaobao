package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderEpocTransfer 电子保单受保人转移
// tmall.servicecenter.spserviceorder.epoc.transfer
//
// 电子保单受保人转移
func TmallServicecenterSpserviceorderEpocTransfer(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderEpocTransferAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderEpocTransferAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
