package aiar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// AlibabaAiArServiceDetect ailab AR图像检索
// alibaba.ai.ar.service.detect
//
// ailab AR图像检索
func AlibabaAiArServiceDetect(ctx context.Context, clt *core.SDKClient, req *aiar.AlibabaAiArServiceDetectAPIRequest, resp *aiar.AlibabaAiArServiceDetectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
