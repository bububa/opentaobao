package aiar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aiar"
)

// Alibabaaiaropenplatformdetect AR开发者平台marker图片检测服务
// alibaba.ai.ar.open.platform.detect
//
// AR开发者平台marker图片检测服务，给AR SDK 和 阿里火眼app使用
func Alibabaaiaropenplatformdetect(clt *core.SDKClient, req *aiar.AlibabaaiaropenplatformdetectAPIRequest, session string) (*aiar.AlibabaaiaropenplatformdetectAPIResponse, error) {
	var resp aiar.AlibabaaiaropenplatformdetectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
