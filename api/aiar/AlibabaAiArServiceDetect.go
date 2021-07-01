package aiar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

/* AlibabaAiArServiceDetect
ailab AR图像检索
alibaba.ai.ar.service.detect

ailab AR图像检索 */
func AlibabaAiArServiceDetect(clt *core.SDKClient, req *aiar.AlibabaAiArServiceDetectAPIRequest, session string) (*aiar.AlibabaAiArServiceDetectAPIResponse, error) {
	var resp aiar.AlibabaAiArServiceDetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
