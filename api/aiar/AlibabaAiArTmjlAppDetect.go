package aiar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// Alibabaaiartmjlappdetect 天猫精灵扫一扫入口的服务
// alibaba.ai.ar.tmjl.app.detect
//
// 天猫精灵扫一扫入口的图像检测服务
func Alibabaaiartmjlappdetect(clt *core.SDKClient, req *aiar.AlibabaaiartmjlappdetectAPIRequest, session string) (*aiar.AlibabaaiartmjlappdetectAPIResponse, error) {
	var resp aiar.AlibabaaiartmjlappdetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
