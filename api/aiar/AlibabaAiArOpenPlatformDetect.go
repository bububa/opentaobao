package aiar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

/* AlibabaAiArOpenPlatformDetect
AR开发者平台marker图片检测服务
alibaba.ai.ar.open.platform.detect

AR开发者平台marker图片检测服务，给AR SDK 和 阿里火眼app使用 */
func AlibabaAiArOpenPlatformDetect(clt *core.SDKClient, req *aiar.AlibabaAiArOpenPlatformDetectAPIRequest, session string) (*aiar.AlibabaAiArOpenPlatformDetectAPIResponse, error) {
	var resp aiar.AlibabaAiArOpenPlatformDetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
