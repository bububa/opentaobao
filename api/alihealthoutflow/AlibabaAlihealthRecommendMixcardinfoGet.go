package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRecommendMixcardinfoGet 快应用混合卡片信息
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
func AlibabaAlihealthRecommendMixcardinfoGet(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRecommendMixcardinfoGetAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthRecommendMixcardinfoGetAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthRecommendMixcardinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
