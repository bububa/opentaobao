package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRecommendMixcardinfoGet 快应用混合卡片信息
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
func AlibabaAlihealthRecommendMixcardinfoGet(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRecommendMixcardinfoGetAPIRequest, resp *alihealthoutflow.AlibabaAlihealthRecommendMixcardinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
