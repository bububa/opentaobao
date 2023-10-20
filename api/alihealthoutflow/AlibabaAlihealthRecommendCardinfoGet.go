package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRecommendCardinfoGet 快应用卡片信息
// alibaba.alihealth.recommend.cardinfo.get
//
// 快应用卡片信息
func AlibabaAlihealthRecommendCardinfoGet(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRecommendCardinfoGetAPIRequest, resp *alihealthoutflow.AlibabaAlihealthRecommendCardinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
