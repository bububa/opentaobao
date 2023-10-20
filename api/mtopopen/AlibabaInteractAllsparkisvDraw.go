package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractAllsparkisvDraw allspark提供抽奖tida接口对应鉴权接口
// alibaba.interact.allsparkisv.draw
//
// 该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
func AlibabaInteractAllsparkisvDraw(clt *core.SDKClient, req *mtopopen.AlibabaInteractAllsparkisvDrawAPIRequest, session string) (*mtopopen.AlibabaInteractAllsparkisvDrawAPIResponse, error) {
	var resp mtopopen.AlibabaInteractAllsparkisvDrawAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
