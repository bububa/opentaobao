package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthrecommendcardinfoget 快应用卡片信息
// alibaba.alihealth.recommend.cardinfo.get
//
// 快应用卡片信息
func Alibabaalihealthrecommendcardinfoget(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthrecommendcardinfogetAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthrecommendcardinfogetAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthrecommendcardinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
