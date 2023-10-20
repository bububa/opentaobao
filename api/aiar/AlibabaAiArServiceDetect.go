package aiar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// AlibabaAiArServiceDetect ailab AR图像检索
// alibaba.ai.ar.service.detect
//
// ailab AR图像检索
func AlibabaAiArServiceDetect(clt *core.SDKClient, req *aiar.AlibabaAiArServiceDetectAPIRequest, resp *aiar.AlibabaAiArServiceDetectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
