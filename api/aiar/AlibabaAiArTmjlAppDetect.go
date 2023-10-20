package aiar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// AlibabaAiArTmjlAppDetect 天猫精灵扫一扫入口的服务
// alibaba.ai.ar.tmjl.app.detect
//
// 天猫精灵扫一扫入口的图像检测服务
func AlibabaAiArTmjlAppDetect(clt *core.SDKClient, req *aiar.AlibabaAiArTmjlAppDetectAPIRequest, session string) (*aiar.AlibabaAiArTmjlAppDetectAPIResponse, error) {
	var resp aiar.AlibabaAiArTmjlAppDetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
