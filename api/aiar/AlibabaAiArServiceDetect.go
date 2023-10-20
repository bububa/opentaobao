package aiar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// Alibabaaiarservicedetect ailab AR图像检索
// alibaba.ai.ar.service.detect
//
// ailab AR图像检索
func Alibabaaiarservicedetect(clt *core.SDKClient, req *aiar.AlibabaaiarservicedetectAPIRequest, session string) (*aiar.AlibabaaiarservicedetectAPIResponse, error) {
	var resp aiar.AlibabaaiarservicedetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
