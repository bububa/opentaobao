package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthrecommendmixcardinfoget 快应用混合卡片信息
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
func Alibabaalihealthrecommendmixcardinfoget(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthrecommendmixcardinfogetAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthrecommendmixcardinfogetAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthrecommendmixcardinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
